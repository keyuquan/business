package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	cfg *Config
)

type Config struct {
	Env     string   `json:"env"`
	Mysql   *Mysql   `json:"mysql"`
	Doris   *Mysql   `json:"doris"`
	Etcd    *Etcd    `json:"etcd"`
	Agora   *Agora   `json:"agora"`
	Redis   *Redis   `json:"redis"`
	Kafka   *Kafka   `json:"kafka"`
	Elastic *Elastic `json:"elastic"`
	Ip      *Ip      `json:"ip"`
}
type Ip struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

func InitConfig() {
	if len(os.Args) <= 1 {
		log.Fatalln("Please specify the configuration file path")
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
	}
	cfg = new(Config)
	err = json.Unmarshal(data, cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("配置文件加载成功")
	//log.Printf("Kafka-broker:%+v \n", cfg.Kafka.Brokers)
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

type Doris struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

func (m *Mysql) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Db)
}

func (m *Mysql) GetDorisDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=False&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Db)
}

type Etcd struct {
	Endpoints []string `json:"endpoints"`
}

type Agora struct {
	Api struct {
		URL    string `json:"url"`
		Key    string `json:"key"`
		Secret string `json:"secret"`
	} `json:"api"`
	AppId       string `json:"app_id"`
	Certificate string `json:"certificate"`
	ExpireTime  uint32 `json:"expire_time"` // 过期时间 单位: d
}

func (a Agora) GetChannelOnlineURL() string {
	return a.Api.URL + "/dev/v1/channel/" + a.AppId
}

func (a Agora) GetChannelOnlineListByChannelId(chanId string) string {
	return a.Api.URL + "/dev/v1/channel/user/" + a.AppId + "/" + chanId
}

func (a Agora) GetSendMsgSignURL() string {
	return a.Api.URL + "/dev/v2/project/" + a.AppId + fmt.Sprintf("/rtm/users/%d/peer_messages", 90909090909090)
}

func (a Agora) GetSendMsgBroadcastURL() string {
	return a.Api.URL + "/dev/v2/project/" + a.AppId + fmt.Sprintf("/rtm/users/%d/channel_messages", 90909090909090)
}

func (a Agora) GetBanPermissionURL() string {
	return a.Api.URL + "/dev/v1/kicking-rule"
}

func (a Agora) GetUserOnlineStatusURL(uid uint64, channelName string) string {
	return a.Api.URL + "/dev/v1/channel/user/property/" + a.AppId + fmt.Sprintf("/%d/%s", uid, channelName)
}

func (a Agora) GetKeyAndSecret() []byte {
	return []byte(a.Api.Key + ":" + a.Api.Secret)
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Auth     string `json:"auth"`
	DbIndex  int    `json:"db_index"`
	Timeout  int    `json:"timeout"`
	PoolSize int    `json:"pool_size"`
}

type Kafka struct {
	Brokers []string `json:"brokers"`
}

type Elastic struct {
	Hosts    []string `json:"hosts"`
	User     string   `json:"user"`
	Password string   `json:"password"`
}

func GetMysql() *Mysql {
	return cfg.Mysql
}

func GetDoris() *Mysql {
	return cfg.Doris
}

func GetEtcd() *Etcd {
	return cfg.Etcd
}

func GetAgora() *Agora {
	return cfg.Agora
}

func GetRedis() *Redis {
	return cfg.Redis
}

func GetEnv() string {
	return cfg.Env
}

func EnvDebug() bool {
	return cfg.Env == "debug"
}

func GetKafkaBroker() []string {
	return cfg.Kafka.Brokers
}

func GetEs() *Elastic {
	return cfg.Elastic
}
func GetIp() *Ip {
	return cfg.Ip
}
