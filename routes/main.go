package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lealhugui/query-reducer/config"
	"github.com/prometheus/common/log"
)

//StartRouter starts the main route server
func StartRouter(cfg config.ServerConfig) {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Info(fmt.Sprintf("Starting Server on Port:%d", cfg.Port))

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)

}
