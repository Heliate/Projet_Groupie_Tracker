package routes

import (
	"Api/controllers"
	"net/http"
)

func SetsPokemonRoutes() {

	http.HandleFunc("/sets_pokemon/", controllers.SetPokemonControllers)
}
