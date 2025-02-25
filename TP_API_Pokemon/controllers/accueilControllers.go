package controllers

import (
	"Api/templates"
	"net/http"
)

func AccueilControllers(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}
