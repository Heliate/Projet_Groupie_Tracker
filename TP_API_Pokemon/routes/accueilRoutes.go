package routes

import (
	"Api/controllers"
	"net/http"
)

func accueilRoutes() {

	http.HandleFunc("/", controllers.AccueilControllers)
}
