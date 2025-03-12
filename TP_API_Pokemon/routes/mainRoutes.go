package routes

import (
	"fmt"
	"net/http"
)

func MainRoutes() {
	// SetsPokemonRoutes()
	cardsRoutes()
	accueilRoutes()
	cguRoutes()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	fmt.Println("http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
