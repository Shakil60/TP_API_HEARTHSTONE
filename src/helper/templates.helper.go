package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Variable globale qui contiendra tous les templates chargés
var listeTemplate *template.Template

// Load charge tous les fichiers HTML depuis le dossier ../templates
func Load() {
	temp, tempErr := template.ParseGlob("../templates/*.html")
	if tempErr != nil {
		log.Fatalf("Erreur template - %s", tempErr.Error())
		return
	}

	listeTemplate = temp
	fmt.Println("Template - chargement des templates terminé")
}

// RenderTemplate exécute le template spécifié
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer

	errRender := listeTemplate.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	buffer.WriteTo(w)
}
