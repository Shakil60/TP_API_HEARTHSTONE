package routes

import (
	"guide/controllers"
	"net/http"
)

// Router enregistre la route /error sur le mux passé par le routeur racine.
// Bonne pratique : on "branche" les routes dans le routeur principal !
func errorRouter(router *http.ServeMux) {
	// Liaison de la route /error au gestionnaire errorDisplay
	router.HandleFunc("/error", controllers.ErrorDisplay)
}