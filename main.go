package main

import (
	"github.com/lealhugui/query-reducer/config"
	"github.com/lealhugui/query-reducer/routes"
)

func main() {
	cfg := config.GetConfig()
	routes.StartRouter(cfg.Server)

}
