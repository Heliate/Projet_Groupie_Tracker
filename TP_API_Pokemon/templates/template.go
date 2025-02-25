package templates

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplates() {
	temp, tempErr := template.ParseGlob("./templates/*.html")
	if tempErr != nil {
		fmt.Printf("Erreur lors du chargement des template : %s\n", tempErr.Error())
		os.Exit(02)
	}
	Temp = temp
}
