package routes

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lealhugui/query-reducer/config"
	"github.com/prometheus/common/log"
)

//StartRouter starts the main route server
func StartRouter(cfg config.ServerConfig) {

	log.Info(fmt.Sprintf("Starting Server..."))

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./static/", false)))
	r.POST("/query", QueryHandler)

	go func() {
		log.Error(r.Run(fmt.Sprintf(":%d", cfg.Port)))
	}()

	log.Info(fmt.Sprintf("Server Started on Port:%d", cfg.Port))

}
