package controllers

import (
	"Api/services/Cartes_Pokemon"
	"Api/templates"
	"net/http"
)

type PageCardsBySet struct {
	NameSet string
	Cards   []Cartes_Pokemon.CartesPokemon
}

// Récupère toutes les cartes d'un seul set
func CardsControllers(w http.ResponseWriter, r *http.Request) {
	nameSet := r.FormValue("set")
	if nameSet == "" {
		nameSet = "swsh1"
	}
	listCards := Cartes_Pokemon.CardsBySetRequest(nameSet)
	if len(listCards) == 0 {
		http.Error(w, "Aucune carte trouvée", http.StatusNotFound)
		return
	}
	data := PageCardsBySet{nameSet, listCards}
	templates.Temp.ExecuteTemplate(w, "cardsBySet", data)
}
