package routes

import (
	"Api/controllers"
	"net/http"
)

func cguRoutes() {
	http.HandleFunc("/cgu", controllers.CguControllers)
}
