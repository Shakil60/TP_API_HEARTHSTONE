package routes

import (
	"net/http"
)

// MainRouter initialise et retourne le routeur principal de l'application
func MainRouter() *http.ServeMux {

	mainRouter := http.NewServeMux()

	errorRouter(mainRouter)

	routerListCards(mainRouter)

	// Configuration du serveur de fichiers statiques (CSS, images, etc.)
	// On sert les assets du projet courant : ../assets (par rapport à src)
	fileServerHandler := http.FileServer(http.Dir("../assets"))

	// Route permettant de servir les fichiers statiques via /static/
	mainRouter.Handle("/static/", http.StripPrefix("/static/", fileServerHandler))

	return mainRouter
}

