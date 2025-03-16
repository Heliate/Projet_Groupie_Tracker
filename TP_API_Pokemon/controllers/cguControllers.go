package controllers

import (
	"Api/templates"
	"net/http"
)

// Contrôleur qui gère l'affichage de la page des CGU (Conditions Générales d'Utilisation)
func CguControllers(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "cgu", nil)
}
