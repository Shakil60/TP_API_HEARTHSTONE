package routes

import (
	"guide/controllers"
	"net/http"
)

func errorRouter(router *http.ServeMux) {
	router.HandleFunc("/error", controllers.ErrorDisplay)
}