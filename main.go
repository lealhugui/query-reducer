package main

import (
	"time"

	"github.com/lealhugui/query-reducer/config"
	"github.com/lealhugui/query-reducer/db"
	"github.com/lealhugui/query-reducer/routes"
	"github.com/prometheus/common/log"
)

var cfg config.AppConfig

func initCfg() {
	cfg = config.GetConfig()
}

func initDb() {
	log.Infof("Starting dbInstance...")
	_, err := db.New(cfg.DbConfigs)
	if err != nil {
		panic(err)
	}
	log.Infof("dbInstance started...")
}

func initServer() {
	routes.StartRouter(cfg.Server)
}

func init() {
	initCfg()
}

func main() {
	initDb()
	initServer()
	//wait forever
	for {
		time.Sleep(100 * time.Millisecond)
	}

}
