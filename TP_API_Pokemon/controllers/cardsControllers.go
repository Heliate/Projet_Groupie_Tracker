package controllers

import (
	"Api/Services/Cartes_Pokemon"
	"Api/templates"
	"Api/utils"
	"net/http"
	"strconv"
)

// Structure pour étendre CartesPokemon avec l'information de favoris
type CardPokemonWithFavorite struct {
	Cartes_Pokemon.CartesPokemon
	IsFavorite bool
}

type PageCardsBySet struct {
	NameSet         string
	Cards           []CardPokemonWithFavorite // Modifié pour utiliser CardPokemonWithFavorite
	TotalPages      int
	CurrentPage     int
	NoCardFound     bool
	NoSearchResults bool
	SearchTerm      string
}

// Récupère toutes les cartes d'un seul set
func CardsControllers(w http.ResponseWriter, r *http.Request) {
	// Dans un environnement réel, vous récupéreriez l'ID utilisateur depuis la session
	userID := "user123"

	nameSet := r.FormValue("set")
	if nameSet == "" {
		nameSet = "swsh1"
	}

	// Récupérer le terme de recherche s'il existe
	searchTerm := r.URL.Query().Get("search")

	listCards := Cartes_Pokemon.CardsBySetRequest(nameSet)
	if len(listCards) == 0 {
		http.Error(w, "Aucune carte trouvée", http.StatusNotFound)
		return
	}

	// Initialiser les flags pour les messages d'erreur
	noSearchResults := false
	noCardFound := false

	// Appliquer la recherche par nom si un terme de recherche est fourni
	filteredBySearch := listCards
	if searchTerm != "" {
		filteredBySearch = utils.SearchCardsByName(listCards, searchTerm)
		// Vérifier si la recherche n'a donné aucun résultat
		if len(filteredBySearch) == 0 {
			noSearchResults = true
			// Si pas de résultats de recherche, on continue avec une liste vide
		}
	}

	page := r.URL.Query().Get("page")
	value, err := strconv.Atoi(page)
	if err != nil || value < 1 {
		value = 1
	}

	filter := utils.GetcardFilters(r)

	// Vérifier si des filtres ont été appliqués
	hasFilters := len(filter.Rarity) > 0 || len(filter.Type) > 0 || len(filter.Weakness) > 0 || filter.SortBy != ""

	// Appliquer les filtres après la recherche (seulement si la recherche a donné des résultats)
	filterCard := filteredBySearch
	if !noSearchResults && hasFilters {
		filterCard = utils.ApplycardFilters(filter, filteredBySearch)
		// Si aucune carte ne correspond aux filtres après une recherche réussie
		if len(filterCard) == 0 {
			noCardFound = true
		}
	}

	cards, totalPages, pages := utils.PaginateCards(filterCard, value, 30)

	// Vérifier si le gestionnaire de favoris est initialisé
	favManager := utils.GetFavoritesManager()

	// Convertir les cartes en cartes avec statut de favoris
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
