package utils

import (
	"Api/Services/Cartes_Pokemon"
	"strings"
)

// Recherche les cartes Pokémon dont le nom contient le terme de recherche
func SearchCardsByName(cards []Cartes_Pokemon.CartesPokemon, searchTerm string) []Cartes_Pokemon.CartesPokemon {
	// Si le terme de recherche est vide, retourne toutes les cartes
	if searchTerm == "" {
		return cards
	}

	// Convertit le terme en minuscules pour une recherche insensible à la casse
	searchTermLower := strings.ToLower(searchTerm)

	// Initialise le tableau de résultats
	searchResults := []Cartes_Pokemon.CartesPokemon{}

	// Parcourt toutes les cartes pour trouver celles correspondant au critère
	for _, card := range cards {
		if strings.Contains(strings.ToLower(card.Name), searchTermLower) {
			searchResults = append(searchResults, card)
		}
	}

	return searchResults
}
