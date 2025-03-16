package routes

import (
	"Api/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func MainRoutes() {
	// Créer le dossier data s'il n'existe pas
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("Erreur lors de la création du dossier de données: %v\n", err)
	}

	// Initialiser le gestionnaire de favoris
	favoritesFile := filepath.Join(dataDir, "favorites.json")
	utils.InitFavoritesManager(favoritesFile)

	cardsRoutes()
	accueilRoutes()
	cguRoutes()
	favorisRoutes()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	fmt.Println("http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
