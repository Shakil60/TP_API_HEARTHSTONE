package routes

import (
	"net/http"
)

// MainRouter initialise et retourne le routeur principal de l'application
func MainRouter() *http.ServeMux {

	// Création du routeur principal
	mainRouter := http.NewServeMux()

	// Enregistrement des routes liées à la gestion des erreurs
	errorRouter(mainRouter)

	// Enregistrement des routes liées aux Cartes
	routerListCards(mainRouter)

	// Configuration du serveur de fichiers statiques (CSS, images, etc.)
	fileServerHandler := http.FileServer(http.Dir("assets"))

	// Route permettant de servir les fichiers statiques via /static/
	mainRouter.Handle("/static/",http.StripPrefix("/static/", fileServerHandler))

	return mainRouter
}

