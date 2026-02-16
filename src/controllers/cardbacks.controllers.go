package controllers

import (
	"guide/helper"
	"guide/services"
	"net/http"
	"strconv"
)

type CardBacksPageData struct {
	CardBacks   []services.CardBack
	CurrentPage int
	TotalPages  int
	HasPrev     bool
	HasNext     bool
	PrevPage    int
	NextPage    int
}

// ListCardsbacksDisplay est un contrôleur permettant de récupérer la liste des dos de cartes.
// Pour l'instant, on renvoie simplement la réponse JSON de l'API.
func ListCardsbacksDisplay(w http.ResponseWriter, r *http.Request) {
	const pageSize = 20
	pageParam := r.URL.Query().Get("page")
	page := 1
	if pageParam != "" {
		if p, errConv := strconv.Atoi(pageParam); errConv == nil && p > 0 {
			page = p
		}
	}

	data, statusCode, err := services.GetCardBacksPage(page, pageSize)
	if err != nil || statusCode != http.StatusOK {
		// S'assurer d'utiliser un code de statut valide
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, "Erreur lors de la récupération des dos de cartes", statusCode)
		return
	}

	// Si aucune donnée n'est renvoyée, on affiche une page vide cohérente
	if data == nil || len(data.CardBacks) == 0 {
		helper.RenderTemplate(w, r, "list_cardbacks", CardBacksPageData{
			CardBacks:   []services.CardBack{},
			CurrentPage: page,
			TotalPages:  1,
			HasPrev:     page > 1,
			HasNext:     false,
			PrevPage:    page - 1,
			NextPage:    page + 1,
		})
		return
	}

	pageData := CardBacksPageData{
		CardBacks:   data.CardBacks,
		CurrentPage: data.Page,
		TotalPages:  data.PageCount,
		HasPrev:     data.Page > 1,
		HasNext:     data.Page < data.PageCount,
		PrevPage:    data.Page - 1,
		NextPage:    data.Page + 1,
	}

	// Chargement du template HTML (grâce au helper)
	helper.RenderTemplate(w, r, "list_cardbacks", pageData)
}
