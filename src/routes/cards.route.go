package routes

import (
	"guide/controllers"
	"net/http"
)

// La fonction routerListCards permet de déclarer l’ensemble des routes liées aux fonctionnalités Cartes.
func routerListCards(router *http.ServeMux) {
	// Déclaration de la route racine "/"
	router.HandleFunc("/", controllers.HomeDisplay)

	// Route pour la page détail d'une carte
	router.HandleFunc("/card", controllers.CardDetailDisplay)

	// Route pour la page de liste des cartes (avec filtres)
	router.HandleFunc("/cards", controllers.ListCardsDisplay)
	
	// Route pour la page de liste des dos de cartes
	router.HandleFunc("/cardbacks", controllers.ListCardsbacksDisplay)

	// Favoris : page liste et ajout
	router.HandleFunc("/favorites", controllers.ListFavoritesDisplay)
	router.HandleFunc("/favorites/add", controllers.AddFavoriteHandler)
	router.HandleFunc("/favorites/remove", controllers.RemoveFavoriteHandler)

	// Route pour la page de saisie du code de deck
	router.HandleFunc("/deck/code", controllers.DeckFormDisplay)

	// Route pour la page d'affichage de la liste des cartes du deck
	router.HandleFunc("/deck", controllers.DeckDisplay)
}
