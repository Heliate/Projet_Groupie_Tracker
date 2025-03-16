package controllers

import (
	"Api/templates"
	"net/http"
)

// Contrôleur qui gère l'affichage de la page d'accueil
func AccueilControllers(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}
