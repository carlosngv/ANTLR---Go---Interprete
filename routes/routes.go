package routes

import (
	"Parser/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func UseRoutes(router *mux.Router) {
	router.HandleFunc("/api/parse", controller.ProcessData).Methods(http.MethodPost)
}
