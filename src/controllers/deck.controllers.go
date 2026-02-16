package controllers

import (
	"guide/helper"
	"guide/services"
	"net/http"
)

// DeckPageData contient les informations nécessaires pour afficher un deck.
type DeckPageData struct {
	DeckCode string
	Deck     *services.Deck
}

// DeckFormData contient les données nécessaires pour la page de saisie du code de deck.
type DeckFormData struct {
	DefaultCode string
}

// DeckFormDisplay affiche une page avec un formulaire permettant de saisir un code de deck.
func DeckFormDisplay(w http.ResponseWriter, r *http.Request) {
	data := DeckFormData{
		DefaultCode: "",
	}
	helper.RenderTemplate(w, r, "deck_form", data)
}

// DeckDisplay est un contrôleur permettant d'afficher un deck à partir d'un code de deck.
// Le code du deck est passé en paramètre de requête : ?code=<deckCode>
func DeckDisplay(w http.ResponseWriter, r *http.Request) {
	deckCode := r.URL.Query().Get("code")
	if deckCode == "" {
		http.Error(w, "Aucun code de deck fourni", http.StatusBadRequest)
		return
	}

	deck, statusCode, err := services.GetDeckPage(deckCode)
	if err != nil || statusCode != http.StatusOK {
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, "Erreur lors de la récupération du deck", statusCode)
		return
	}

	pageData := DeckPageData{
		DeckCode: deckCode,
		Deck:     deck,
	}

	// À adapter au nom réel de votre template (par ex. "deck" ou "deck_details")
	helper.RenderTemplate(w, r, "deck", pageData)
}
