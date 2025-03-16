package routes

import (
	"Api/controllers"
	"net/http"
)

// Configure la route pour la page d'accueil
func accueilRoutes() {
	http.HandleFunc("/", controllers.AccueilControllers)
}
