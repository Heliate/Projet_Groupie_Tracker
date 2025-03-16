package utils

import (
	"Api/Services/Cartes_Pokemon"
	"math"
)

// Fonction qui divise une liste de cartes en pages
// Retourne la page demandée, le nombre total de pages, et le numéro de la page courante
func PaginateCards(data []Cartes_Pokemon.CartesPokemon, page int, cardsPerPages int) ([]Cartes_Pokemon.CartesPokemon, int, int) {
	// Gestion du cas d'une liste vide
	if len(data) == 0 {
		return []Cartes_Pokemon.CartesPokemon{}, 0, 1
	}

	// Calcul du nombre total de pages
	TotalCards := len(data)
	TotalPages := int(math.Ceil(float64(TotalCards) / float64(cardsPerPages)))

	// Vérifie que le numéro de page est valide
	if page > TotalPages {
		page = TotalPages
	}
	if page < 1 {
		page = 1
	}

	// Calcul des indices de début et de fin pour le découpage
	start := (page - 1) * cardsPerPages
	end := start + cardsPerPages

	// Ajustement des indices si nécessaire
	if start > TotalCards {
		start = TotalCards
	}
	if end > TotalCards {
		end = TotalCards
	}

	// Retourne la tranche correspondant à la page demandée
	return data[start:end], TotalPages, page
}
