package db

import (
	"business/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/olivere/elastic/v7"
	etcdclient "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db       *gorm.DB
	dorisDb  *gorm.DB
	etcd     *etcdclient.Client
	cache    *redis.Client
	esCli    *elastic.Client
	ipSearch *xdb.Searcher
)

func InitMysql(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute)
	if err = sqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Mysql加载成功")
}

func InitDoris(dsn string) {
	var err error
	dorisDb, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB, err := dorisDb.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(2)
	sqlDB.SetConnMaxLifetime(time.Minute)
	if err = sqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}
	log.Println("doris加载成功")
}

func GetDorisDB() *gorm.DB {
	if config.GetEnv() == "debug" {
		return dorisDb.Debug()
	}
	return dorisDb
}

func GetDB() *gorm.DB {
	if config.GetEnv() == "debug" {
		return db.Debug()
	}
	return db
}

func InitEtcd(endpoints []string) {
	var err error
	etcd, err = etcdclient.New(etcdclient.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatalf("init etcd cli is error: %+v \n", err)
	}
	log.Println("Etcd加载成功")
}

func GetEtcd() *etcdclient.Client {
	return etcd
}

func InitRedis(cfg *config.Redis) {
	cache = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:          cfg.DbIndex,
		Password:    cfg.Auth,
		PoolTimeout: time.Duration(cfg.Timeout) * time.Second,
		PoolSize:    cfg.PoolSize,
	})
	if err := cache.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("init redis cli is error: %+v \n", err.Error())
	}
	log.Println("Redis加载成功")
}

func GetRedis() *redis.Client {
	return cache
}

func RedisLock(key string, timeout time.Duration) (bool, error) {
	return cache.SetNX(context.Background(), key, 1, timeout*time.Second).Result()
}

func Unlock(key string) error {
	return cache.Del(context.Background(), key).Err()
}

func InitEsClient(es *config.Elastic) {
	var err error
	esCli, err = elastic.NewClient(
		elastic.SetURL(es.Hosts...),
		elastic.SetGzip(true),
		elastic.SetHealthcheckTimeout(10*time.Second),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(es.User, es.Password),
	)
	if err != nil {
		log.Fatalln("init es client is error: " + err.Error())
	}
}

func GetEsClient() *elastic.Client {
	return esCli
}
func GetIPSearch() *xdb.Searcher {
	return ipSearch
}
func InitIPParse(dbPath string, cachePolicy string) {
	var (
		err  error
		body []byte
	)
	switch cachePolicy {
	case "nil", "file":
		ipSearch, err = xdb.NewWithFileOnly(dbPath)
	case "vectorIndex":
		body, err = xdb.LoadVectorIndexFromFile(dbPath)
		if err != nil {
			log.Fatalf("init ip path is error: %+v", err)
		}

		ipSearch, err = xdb.NewWithVectorIndex(dbPath, body)
	case "content":
		body, err = xdb.LoadContentFromFile(dbPath)
		if err != nil {
			log.Fatalf("init ip path is error: %+v", err)
		}
		ipSearch, err = xdb.NewWithBuffer(body)
		log.Println("IP库正在加载")
	default:
		log.Fatalln("IP没有该类型")
	}
	if err != nil {
		if err != nil {
			log.Fatalf("load ip xdb is error: %+v", err)
		}
	}
}
