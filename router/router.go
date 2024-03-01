package router

import (
	"cross-check/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/", middleware.Home)

	return router
}
