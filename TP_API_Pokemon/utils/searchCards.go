package utils

import (
	"Api/Services/Cartes_Pokemon"
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
	// Si aucun filtre n'est appliqué, retourner toutes les cartes
	if len(filters.Rarity) == 0 && len(filters.Type) == 0 && len(filters.Weakness) == 0 && filters.SortBy == "" {
		return Allcards
	}

	hasTypeFilters := len(filters.Type) > 0
	hasRarityFilters := len(filters.Rarity) > 0
	hasWeaknessFilters := len(filters.Weakness) > 0

	// Appliquer les filtres un par un
	var filteredCards []Cartes_Pokemon.CartesPokemon

	// Commence avec toutes les cartes
	filteredCards = Allcards

	// Appliquer les filtres de rareté si présents
	if hasRarityFilters {
		filteredCards = filtercardsByRarity(filteredCards, filters.Rarity)
	}

	// Appliquer les filtres de type si présents
	if hasTypeFilters {
		filteredCards = FiltercardsByTypes(filteredCards, filters.Type)
	}

	// Appliquer les filtres de faiblesse si présents
	if hasWeaknessFilters {
		filteredCards = FiltercardsByWeakness(filteredCards, filters.Weakness)
	}

	// Tri des cartes si demandé et si des cartes existent
	if filters.SortBy != "" && len(filteredCards) > 0 {
		switch filters.SortBy {
		case "0":
			filteredCards = filtercardsByHPCresc(filteredCards)
		case "1":
			filteredCards = filtercardsByHPDecresc(filteredCards)
		}
	}

	return filteredCards
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

	for _, card := range cards {
		for _, t := range types {
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
