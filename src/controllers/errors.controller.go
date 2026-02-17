package controllers

import (
	"net/http"
	"guide/services"
	"guide/helper"
)

// errorDisplay gère l'affichage de la page d'erreur.
// On récupère le code et le message passés en query string /error?code=404&message=Page%30introuvable
func ErrorDisplay(w http.ResponseWriter, r *http.Request) {
	data := services.Error{
		Code:    r.FormValue("code"),    // "404"
		Message: r.FormValue("message"), // "Page introuvable"
	}

	helper.RenderTemplate(w, r, "error", data)
}
