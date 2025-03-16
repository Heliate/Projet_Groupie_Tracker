package controllers

import (
	"Api/Services/Cartes_Pokemon"
	"Api/templates"
	"Api/utils"
	"net/http"
	"strconv"
)

type CardPokemonWithFavorite struct {
	Cartes_Pokemon.CartesPokemon
	IsFavorite bool
}

type PageCardsBySet struct {
	NameSet         string
	Cards           []CardPokemonWithFavorite
	TotalPages      int
	CurrentPage     int
	NoCardFound     bool
	NoSearchResults bool
	SearchTerm      string
}

// Contrôleur qui gère l'affichage des cartes d'un set avec filtrage, recherche, pagination et état des favoris
func CardsControllers(w http.ResponseWriter, r *http.Request) {
	// Simuler un ID utilisateur
	userID := "user123"

	// Récupérer le set demandé ou utiliser le set par défaut
	nameSet := r.FormValue("set")
	if nameSet == "" {
		nameSet = "swsh1"
	}

	// Récupérer le terme de recherche
	searchTerm := r.URL.Query().Get("search")

	// Récupérer toutes les cartes du set
	listCards := Cartes_Pokemon.CardsBySetRequest(nameSet)
	if len(listCards) == 0 {
		http.Error(w, "Aucune carte trouvée", http.StatusNotFound)
		return
	}

	// Initialiser les flags pour les messages
	noSearchResults := false
	noCardFound := false

	// Filtrer par recherche si terme fourni
	filteredBySearch := listCards
	if searchTerm != "" {
		filteredBySearch = utils.SearchCardsByName(listCards, searchTerm)
		if len(filteredBySearch) == 0 {
			noSearchResults = true
		}
	}

	// Récupérer la page demandée
	page := r.URL.Query().Get("page")
	value, err := strconv.Atoi(page)
	if err != nil || value < 1 {
		value = 1
	}

	// Récupérer et appliquer les filtres
	filter := utils.GetcardFilters(r)
	hasFilters := len(filter.Rarity) > 0 || len(filter.Type) > 0 || len(filter.Weakness) > 0 || filter.SortBy != ""

	filterCard := filteredBySearch
	if !noSearchResults && hasFilters {
		filterCard = utils.ApplycardFilters(filter, filteredBySearch)
		if len(filterCard) == 0 {
			noCardFound = true
		}
	}

	// Paginer les résultats
	cards, totalPages, pages := utils.PaginateCards(filterCard, value, 30)

	// Récupérer l'état des favoris pour chaque carte
	favManager := utils.GetFavoritesManager()
	cardsWithFavorite := make([]CardPokemonWithFavorite, len(cards))
	for i, card := range cards {
		isFavorite := false
		if favManager != nil {
			isFavorite, _ = favManager.IsCardInFavorites(userID, card.ID)
		}
		cardsWithFavorite[i] = CardPokemonWithFavorite{
			CartesPokemon: card,
			IsFavorite:    isFavorite,
		}
	}

	// Préparer les données et afficher le template
	data := PageCardsBySet{
		NameSet:         nameSet,
		Cards:           cardsWithFavorite,
		TotalPages:      totalPages,
		CurrentPage:     pages,
		NoCardFound:     noCardFound,
		NoSearchResults: noSearchResults,
		SearchTerm:      searchTerm,
	}

	templates.Temp.ExecuteTemplate(w, "cardsBySet", data)
}
