package controllers

import (
	"Api/services/Cartes_Pokemon"
	"Api/templates"
	"Api/utils"
	"net/http"
	"strconv"
)

type PageCardsBySet struct {
	NameSet     string
	Cards       []Cartes_Pokemon.CartesPokemon
	TotalPages  int
	CurrentPage int
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

	page := r.URL.Query().Get("page")
	value, err := strconv.Atoi(page)
	if err != nil || value < 1 {
		value = 1
	}

	filter := utils.GetcardFilters(r)

	filterCard := utils.ApplycardFilters(filter, listCards)
	//condition à rajouter pour message d'erreur en cas de cartes ne correspondant pas aux filtres

	cards, totalPages, pages := utils.PaginateCards(filterCard, value, 30)

	data := PageCardsBySet{nameSet, cards, totalPages, pages}
	templates.Temp.ExecuteTemplate(w, "cardsBySet", data)
}
