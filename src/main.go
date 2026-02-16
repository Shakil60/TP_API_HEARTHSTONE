package main

import (
	"fmt"
	"guide/helper"
	"guide/routes"
	"guide/services"
	"net/http"
)

func main() {
	helper.Load()

	// Récupération du token OAuth Battle.net (optionnel : utilise les identifiants fournis ou les variables d'environnement)
	clientID := "ba14a7ee2ba2470d86231d4c998043bb"
	clientSecret := "Hz78Ri19bVgSF0TWhl8QnXwbCAK4qst5"
	if token, _, err := services.GetBattleNetToken(clientID, clientSecret); err == nil {
		services.Token = token
		fmt.Println("Token OAuth Battle.net chargé.")
	} else {
		fmt.Printf("Attention: token OAuth non chargé (%v), utilisation du token par défaut.\n", err)
	}

	serveRouter := routes.MainRouter()
	fmt.Println("Serveur lancé : http://localhost:8080")
	http.ListenAndServe("localhost:8080", serveRouter)
}
