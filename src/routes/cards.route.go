package routes

import (
	"guide/controllers"
	"net/http"
)

// La fonction routerListCards permet de déclarer l’ensemble des routes liées aux fonctionnalités Cartes.
// Elle reçoit en paramètre un pointeur vers un http.ServeMux afin d’enregistrer les routes directement sur le routeur principal
func routerListCards(router *http.ServeMux) {
	// Déclaration de la route racine "/"
	router.HandleFunc("/", controllers.HomeDisplay)

	// Route pour la page de liste des cartes (avec filtres)
	router.HandleFunc("/cards", controllers.ListCardsDisplay)

	router.HandleFunc("/cardbacks", controllers.ListCardsbacksDisplay)

	// Favoris : page liste et ajout
	router.HandleFunc("/favorites", controllers.ListFavoritesDisplay)
	router.HandleFunc("/favorites/add", controllers.AddFavoriteHandler)
	router.HandleFunc("/favorites/remove", controllers.RemoveFavoriteHandler)

	// Route pour la page de saisie du code de deck
	router.HandleFunc("/deck/code", controllers.DeckFormDisplay)

	router.HandleFunc("/deck", controllers.DeckDisplay)
}
