package main

import (
	"Api/routes"
	"Api/templates"
)

func main() {
	templates.InitTemplates()
	routes.MainRoutes()
}
