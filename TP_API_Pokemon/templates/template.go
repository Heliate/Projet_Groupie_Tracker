package templates

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func add(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }

var funcMap = template.FuncMap{
	"add": add,
	"sub": sub,
}

func InitTemplates() {
	temp, tempErr := template.New("").Funcs(funcMap).ParseGlob("./templates/*.html")
	if tempErr != nil {
		fmt.Printf("Erreur lors du chargement des template : %s\n", tempErr.Error())
		os.Exit(02)
	}
	Temp = temp
}
