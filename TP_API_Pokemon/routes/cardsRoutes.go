package routes

import (
	"Api/controllers"
	"net/http"
)

func cardsRoutes() {

	http.HandleFunc("/cards", controllers.CardsControllers)
}
