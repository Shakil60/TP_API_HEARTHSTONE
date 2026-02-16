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

// CardsFilterPageData contient les données pour la page liste des cartes avec filtres.
type CardsFilterPageData struct {
	Cards       []services.Card
	CurrentPage int
	TotalPages  int
	HasPrev     bool
	HasNext     bool
	PrevURL     string
	NextURL     string
	ManaCost    string
	Attack      string
	Health      string
	Rarity      string
	SearchTerm  string
}

// ListCardsFilterDisplay affiche la liste des cartes avec filtres ManaCost, Attack, Health, Rarity.
func ListCardsFilterDisplay(w http.ResponseWriter, r *http.Request) {
	const pageSize = 20

	pageParam := r.URL.Query().Get("page")
	page := 1
	if pageParam != "" {
		if p, errConv := strconv.Atoi(pageParam); errConv == nil && p > 0 {
			page = p
		}
	}

	manaCost := strings.TrimSpace(r.URL.Query().Get("manaCost"))
	attack := strings.TrimSpace(r.URL.Query().Get("attack"))
	health := strings.TrimSpace(r.URL.Query().Get("health"))
	rarity := strings.TrimSpace(r.URL.Query().Get("rarity"))
	searchTerm := strings.TrimSpace(r.URL.Query().Get("textFilter"))

	filters := services.CardFilters{
		ManaCost:   manaCost,
		Attack:     attack,
		Health:     health,
		Rarity:     rarity,
		TextFilter: searchTerm,
	}

	data, statusCode, err := services.GetCardsPageWithFilters(page, pageSize, filters)
	if err != nil || statusCode != http.StatusOK {
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des cartes: %v", err), statusCode)
		return
	}

	buildPageURL := func(targetPage int) string {
		if targetPage < 1 {
			return ""
		}
		q := url.Values{}
		q.Set("page", strconv.Itoa(targetPage))
		if manaCost != "" {
			q.Set("manaCost", manaCost)
		}
		if attack != "" {
			q.Set("attack", attack)
		}
		if health != "" {
			q.Set("health", health)
		}
		if rarity != "" {
			q.Set("rarity", rarity)
		}
		if searchTerm != "" {
			q.Set("textFilter", searchTerm)
		}
		return "/cards/filter?" + q.Encode()
	}

	pageData := CardsFilterPageData{
		Cards:       data.Cards,
		CurrentPage: data.Page,
		TotalPages:  data.PageCount,
		HasPrev:     data.Page > 1,
		HasNext:     data.Page < data.PageCount,
		PrevURL:     buildPageURL(data.Page - 1),
		NextURL:     buildPageURL(data.Page + 1),
		ManaCost:    manaCost,
		Attack:      attack,
		Health:      health,
		Rarity:      rarity,
		SearchTerm:  searchTerm,
	}

	// Cas page vide (aucune carte)
	if len(data.Cards) == 0 {
		pageData.Cards = []services.Card{}
		pageData.TotalPages = 1
		pageData.HasNext = false
		pageData.NextURL = ""
	}

	helper.RenderTemplate(w, r, "list_cards_filter", pageData)
}
