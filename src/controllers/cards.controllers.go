package controllers

import (
	"fmt"
	"guide/helper"
	"guide/services"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CardsPageData contient les informations nécessaires pour l'affichage paginé
type CardsPageData struct {
	Cards       []services.Card
	CurrentPage int
	TotalPages  int
	HasPrev     bool
	HasNext     bool
	PrevPage    int
	NextPage    int
	PrevURL     string
	NextURL     string
	SearchTerm  string
}

// ListCardsDisplay est un contrôleur permettant d'afficher la liste des cartes avec pagination
func ListCardsDisplay(w http.ResponseWriter, r *http.Request) {
	// Récupération du numéro de page depuis les paramètres de requête (?page=)
	const pageSize = 20
	pageParam := r.URL.Query().Get("page")
	page := 1
	if pageParam != "" {
		if p, errConv := strconv.Atoi(pageParam); errConv == nil && p > 0 {
			page = p
		}
	}

	// Récupération de la recherche textuelle (texte de carte)
	searchTerm := strings.TrimSpace(r.URL.Query().Get("textFilter"))

	// Appel de la fonction pour récupérer une page de cartes depuis l'API
	data, statusCode, err := services.GetCardsPage(page, pageSize, searchTerm)
	if err != nil || statusCode != http.StatusOK {
		// Utiliser un code d'erreur par défaut si statusCode est 0
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		errorMsg := fmt.Sprintf("Erreur lors de la récupération des Cartes: %v", err)
		http.Error(w, errorMsg, statusCode)
		return
	}

	buildPageURL := func(targetPage int) string {
		if targetPage < 1 {
			return ""
		}
		q := url.Values{}
		q.Set("page", strconv.Itoa(targetPage))
		if searchTerm != "" {
			q.Set("textFilter", searchTerm)
		}
		return "/cards?" + q.Encode()
	}

	// Si aucune carte n'est renvoyée, on affiche une page vide cohérente
	if len(data.Cards) == 0 {
		helper.RenderTemplate(w, r, "list_cards", CardsPageData{
			Cards:       []services.Card{},
			CurrentPage: page,
			TotalPages:  1,
			HasPrev:     page > 1,
			HasNext:     false,
			PrevPage:    page - 1,
			NextPage:    page + 1,
			PrevURL:     buildPageURL(page - 1),
			NextURL:     "",
			SearchTerm:  searchTerm,
		})
		return
	}

	pageData := CardsPageData{
		Cards:       data.Cards,
		CurrentPage: data.Page,
		TotalPages:  data.PageCount,
		HasPrev:     data.Page > 1,
		HasNext:     data.Page < data.PageCount,
		PrevPage:    data.Page - 1,
		NextPage:    data.Page + 1,
		PrevURL:     buildPageURL(data.Page - 1),
		NextURL:     buildPageURL(data.Page + 1),
		SearchTerm:  searchTerm,
	}

	// Chargement du template HTML (grâce au helper)
	helper.RenderTemplate(w, r, "list_cards", pageData)
}