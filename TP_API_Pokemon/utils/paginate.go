package utils

import (
	"Api/services/Cartes_Pokemon"
	"math"
)

func PaginateCards(data []Cartes_Pokemon.CartesPokemon, page int, cardsPerPages int) ([]Cartes_Pokemon.CartesPokemon, int, int) {
	TotalCards := len(data)
	TotalPages := int(math.Ceil(float64(TotalCards) / float64(cardsPerPages)))
	if page > TotalPages {
		page = TotalPages
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
