package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lealhugui/query-reducer/config"
	"github.com/prometheus/common/log"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var url = r.RequestURI
		if url == "/" || url == "" {
			url = "/index.html"
		}

		log.Info(fmt.Sprintf("Server:[%s]::%s", r.Method, url))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

//StartRouter starts the main route server
func StartRouter(cfg config.ServerConfig) {

	log.Info(fmt.Sprintf("Starting Server..."))

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./static/", false)))
	r.POST("/query", QueryHandler)

	/*
		r := mux.NewRouter()
		r.HandleFunc("/query", QueryHandler)
		r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
		r.Use(loggingMiddleware)
	*/

	go func() {
		log.Error(r.Run(fmt.Sprintf(":%d", cfg.Port)))
	}()

	log.Info(fmt.Sprintf("Server Started on Port:%d", cfg.Port))

}
