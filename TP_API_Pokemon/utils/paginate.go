package utils

import (
	"Api/Services/Cartes_Pokemon"
	"math"
)

func PaginateCards(data []Cartes_Pokemon.CartesPokemon, page int, cardsPerPages int) ([]Cartes_Pokemon.CartesPokemon, int, int) {
	// Gestion du cas où le tableau est vide
	if len(data) == 0 {
		return []Cartes_Pokemon.CartesPokemon{}, 0, 1
	}

	TotalCards := len(data)
	TotalPages := int(math.Ceil(float64(TotalCards) / float64(cardsPerPages)))

	// Vérifier que la page est valide
	if page > TotalPages {
		page = TotalPages
	}
	if page < 1 {
		page = 1
	}

	start := (page - 1) * cardsPerPages
	end := start + cardsPerPages

	if start > TotalCards {
		start = TotalCards
	}
	if end > TotalCards {
		end = TotalCards
	}

	return data[start:end], TotalPages, page
}
