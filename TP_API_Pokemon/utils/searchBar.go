package utils

import (
	"Api/Services/Cartes_Pokemon"
	"strings"
)

// SearchCardsByName recherche les cartes Pokémon dont le nom contient le terme de recherche
func SearchCardsByName(cards []Cartes_Pokemon.CartesPokemon, searchTerm string) []Cartes_Pokemon.CartesPokemon {
	// Si le terme de recherche est vide, retourner toutes les cartes
	if searchTerm == "" {
		return cards
	}

	// Convertir le terme de recherche en minuscules pour une recherche insensible à la casse
	searchTermLower := strings.ToLower(searchTerm)

	// Initialiser le slice pour stocker les résultats de la recherche
	searchResults := []Cartes_Pokemon.CartesPokemon{}

	// Parcourir toutes les cartes et vérifier si leur nom contient le terme de recherche
	for _, card := range cards {
		if strings.Contains(strings.ToLower(card.Name), searchTermLower) {
			searchResults = append(searchResults, card)
		}
	}

	return searchResults
}
