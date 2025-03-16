package routes

import (
	"Api/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// Configure toutes les routes de l'application et démarre le serveur
func MainRoutes() {
	// Créer le dossier data pour le stockage
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("Erreur lors de la création du dossier de données: %v\n", err)
	}

	// Initialiser le système de gestion des favoris
	favoritesFile := filepath.Join(dataDir, "favorites.json")
	utils.InitFavoritesManager(favoritesFile)

	// Configurer toutes les routes de l'application
	cardsRoutes()
	accueilRoutes()
	cguRoutes()
	favorisRoutes()

	// Configurer la route pour les fichiers statiques (CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	// Afficher l'URL du serveur
	fmt.Println("http://localhost:8000")

	// Démarrer le serveur sur le port 8000
	http.ListenAndServe(":8000", nil)
}
