package main

import (
	"time"

	"github.com/lealhugui/query-reducer/config"
	"github.com/lealhugui/query-reducer/db"
	"github.com/lealhugui/query-reducer/routes"
	"github.com/prometheus/common/log"
)

func main() {
	cfg := config.GetConfig()
	routes.StartRouter(cfg.Server)

	ctr, err := db.New(cfg.DbConfigs)
	if err != nil {
		log.Error(err)
	} else {
		log.Info(ctr.Query("select * from \"pdv-va\".parametro"))
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}

}
