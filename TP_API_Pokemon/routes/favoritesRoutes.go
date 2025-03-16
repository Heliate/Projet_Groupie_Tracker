package routes

import (
	"Api/controllers"
	"net/http"
)

func favorisRoutes() {
	// Route pour afficher la page des favoris
	http.HandleFunc("/favoris", controllers.FavorisController)

	// Route pour ajouter une carte aux favoris
	http.HandleFunc("/favoris/add", controllers.AddToFavorisController)

	// Route pour supprimer une carte des favoris
	http.HandleFunc("/favoris/remove", controllers.RemoveFromFavorisController)
}
