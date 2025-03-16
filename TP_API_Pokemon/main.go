package main

import (
	"Api/routes"
	"Api/templates"
)

// Point d'entrée de l'application
func main() {
	// Initialise les templates HTML
	templates.InitTemplates()

	// Configure les routes et démarre le serveur
	routes.MainRoutes()
}
