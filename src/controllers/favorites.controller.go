package controllers

import (
	"guide/helper"
	"guide/services"
	"net/http"
	"strconv"
	"strings"
)

// ListFavoritesDisplay affiche la page des cartes favoris.
func ListFavoritesDisplay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	list, err := services.LoadFavorites()
	if err != nil {
		http.Error(w, "Erreur lors du chargement des favoris: "+err.Error(), http.StatusInternalServerError)
		return
	}
	helper.RenderTemplate(w, r, "favorites", list)
}

// AddFavoriteHandler ajoute une carte aux favoris (POST avec formulaire).
func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/favorites", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulaire invalide", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(strings.TrimSpace(r.PostFormValue("id")))
	if id == 0 {
		http.Redirect(w, r, "/cards", http.StatusSeeOther)
		return
	}

	manaCost, _ := strconv.Atoi(r.PostFormValue("manaCost"))
	attack, _ := strconv.Atoi(r.PostFormValue("attack"))
	health, _ := strconv.Atoi(r.PostFormValue("health"))
	rarity, _ := strconv.Atoi(r.PostFormValue("rarity"))
	name := strings.TrimSpace(r.PostFormValue("name"))
	image := strings.TrimSpace(r.PostFormValue("image"))

	card := services.FavoriteCard{
		ID:       id,
		ManaCost: manaCost,
		Attack:   attack,
		Health:   health,
		Rarity:   rarity,
		Name:     name,
		Image:    image,
	}

	_ = services.AddFavorite(card)
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

// RemoveFavoriteHandler retire une carte des favoris (POST avec formulaire).
func RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/favorites", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulaire invalide", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(strings.TrimSpace(r.PostFormValue("id")))
	if id == 0 {
		http.Redirect(w, r, "/favorites", http.StatusSeeOther)
		return
	}

	_ = services.RemoveFavorite(id)
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}
