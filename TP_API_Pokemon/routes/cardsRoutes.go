package routes

import (
	"Api/controllers"
	"net/http"
)

// Configure la route pour l'affichage des cartes
func cardsRoutes() {
	http.HandleFunc("/cards", controllers.CardsControllers)
}
