package templates

import (
	"fmt"
	"html/template"
	"os"
)

// Variable globale pour stocker les templates compilés
var Temp *template.Template

// Fonctions auxiliaires pour les templates
func add(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }

// Map des fonctions disponibles dans les templates
var funcMap = template.FuncMap{
	"add": add,
	"sub": sub,
}

// Initialise tous les templates HTML avec les fonctions personnalisées
func InitTemplates() {
	temp, tempErr := template.New("").Funcs(funcMap).ParseGlob("./templates/*.html")
	if tempErr != nil {
		fmt.Printf("Erreur lors du chargement des template : %s\n", tempErr.Error())
		os.Exit(02)
	}
	Temp = temp
}
