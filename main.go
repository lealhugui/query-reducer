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
	log.Infof("Starting dbInstance...")
	_, err := db.New(cfg.DbConfigs)
	if err != nil {
		panic(err)
	}
	log.Infof("dbInstance started...")
	routes.StartRouter(cfg.Server)

	if err != nil {
		log.Error(err)
	} /*
		else {
			qResult, _ := ctr.Query("select chave, valor from \"pdv-va\".parametro limit 3")
			//fmt.Printf("%# v", pretty.Formatter(aggregator.AggregateResultSet(qResult)))
			encoder.ResultSetEncoder(aggregator.AggregateResultSet(qResult))
		}
	*/

	for {
		time.Sleep(100 * time.Millisecond)
	}

}
