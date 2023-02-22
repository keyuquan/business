package main

import (
	"business/config"
	"business/internal/router"
	"business/pkg/db"
	"business/pkg/log"
	"net/http"
)

func main() {
	log.InitLog()
	// 1. load config
	config.InitConfig()
	// 2. init redis
	//db.InitRedis(config.GetRedis())
	// 3. init mysql
	db.InitMysql(config.GetMysql().GetDsn())
	// 4. init router
	router := router.InitRouter()
	server := &http.Server{
		Addr:    ":10085",
		Handler: router,
	}
	// 4. run http server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server startup err: %s", err.Error())
	}
}
