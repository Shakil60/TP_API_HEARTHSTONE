package controllers

import (
	"guide/helper"
	"net/http"
)

// HomeDisplay est le contrôleur pour la page d'accueil
func HomeDisplay(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, r, "homepage", nil)
}
