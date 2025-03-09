package controllers

import (
	"Api/templates"
	"net/http"
)

func CguControllers(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "cgu", nil)
}
