package routes

import (
	"Api/controllers"
	"net/http"
)

// Configure la route pour la page des CGU
func cguRoutes() {
	http.HandleFunc("/cgu", controllers.CguControllers)
}
