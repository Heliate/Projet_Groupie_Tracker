package utils

import (
	"Api/services/Cartes_Pokemon"
	"net/http"
	"sort"
)

type cardFilters struct {
	Rarity   []string
	Type     []string
	Weakness []string
	SortBy   string
}

func GetcardFilters(r *http.Request) cardFilters {
	filters := cardFilters{
		Rarity:   r.URL.Query()["rarity_filter"],
		Type:     r.URL.Query()["type_filter"],
		Weakness: r.URL.Query()["weakness_filter"],
		SortBy:   r.URL.Query().Get("sort_filter"),
	}
	return filters
}

func ApplycardFilters(filters cardFilters, Allcards []Cartes_Pokemon.CartesPokemon) []Cartes_Pokemon.CartesPokemon {
	cards := Allcards

	// Appliquer tous les filtres sur l'ensemble initial des cartes
	if len(filters.Rarity) > 0 {
		cards = filtercardsByRarity(cards, filters.Rarity)
	}

	if len(filters.Type) > 0 {
		cards = FiltercardsByTypes(cards, filters.Type)
	}

	if len(filters.Weakness) > 0 {
		cards = FiltercardsByWeakness(cards, filters.Weakness)
	}

	// Si aucun filtre n'est appliqué, retourner toutes les cartes
	if len(filters.Rarity) == 0 && len(filters.Type) == 0 && len(filters.Weakness) == 0 {
		cards = Allcards
	}

	if len(cards) == 0 {
		cards = Allcards //aucune cartes n'a trouve dans les filtres
	}

	// Tri des cartes si demandé
	if filters.SortBy != "" {
		switch filters.SortBy {
		case "0":
			cards = filtercardsByHPCresc(cards)
		case "1":
			cards = filtercardsByHPDecresc(cards)
		}
	}

	return cards
}

func filtercardsByRarity(cards []Cartes_Pokemon.CartesPokemon, rarities []string) []Cartes_Pokemon.CartesPokemon {
	filteredcards := []Cartes_Pokemon.CartesPokemon{}
	for _, card := range cards {
		for _, rarity := range rarities {
			if card.Rarity == rarity {
				filteredcards = append(filteredcards, card)
				break
			}
		}
	}
	return filteredcards
}

func FiltercardsByWeakness(cards []Cartes_Pokemon.CartesPokemon, types []string) []Cartes_Pokemon.CartesPokemon {
	filteredcards := []Cartes_Pokemon.CartesPokemon{}
	for _, card := range cards {
		for _, t := range types {
			for _, weakness := range card.Weaknesses {
				if weakness.Type == t {
					filteredcards = append(filteredcards, card)
					break
				}
			}
		}
	}
	return filteredcards
}

func filtercardsByHPCresc(cards []Cartes_Pokemon.CartesPokemon) []Cartes_Pokemon.CartesPokemon {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Hp < cards[j].Hp
	})
	return cards
}

func filtercardsByHPDecresc(cards []Cartes_Pokemon.CartesPokemon) []Cartes_Pokemon.CartesPokemon {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Hp > cards[j].Hp
	})
	return cards
}

func FiltercardsByTypes(cards []Cartes_Pokemon.CartesPokemon, types []string) []Cartes_Pokemon.CartesPokemon {
	filteredcards := []Cartes_Pokemon.CartesPokemon{}

	for _, t := range types {
		for _, card := range cards {
			for _, cardType := range card.Types {
				if cardType == t {
					filteredcards = append(filteredcards, card)
					break
				}
			}
		}
	}
	return filteredcards
}
