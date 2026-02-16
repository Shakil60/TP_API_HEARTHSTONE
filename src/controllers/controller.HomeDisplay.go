package controllers

import (
	"guide/helper"
	"net/http"
)

// HomeDisplay est le contrôleur pour la page d'accueil
func HomeDisplay(w http.ResponseWriter, r *http.Request) {
	// Pas de données spécifiques nécessaires pour la page d'accueil
	helper.RenderTemplate(w, r, "homepage", nil)
}
