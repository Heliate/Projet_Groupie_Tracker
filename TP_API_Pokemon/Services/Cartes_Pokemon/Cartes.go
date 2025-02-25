package Cartes_Pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CartesPokemon struct {
	Category    string   `json:"category"`
	ID          string   `json:"id"`
	Illustrator string   `json:"illustrator"`
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Rarity      string   `json:"rarity"`
	DexID       []int    `json:"dexId"`
	Hp          int      `json:"hp"`
	Types       []string `json:"types"`
	Stage       string   `json:"stage"`
	Attacks     []struct {
		Cost   []string    `json:"cost"`
		Name   string      `json:"name"`
		Effect string      `json:"effect"`
		Damage interface{} `json:"damage"`
	} `json:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
	Retreat int `json:"retreat"`
}

func CardsBySetRequest(name string) []CartesPokemon {
	result := []CartesPokemon{}
	j := 1
	for {
		url := fmt.Sprintf("https://api.tcgdex.net/v2/fr/cards/%s-%d", name, j)

		httpClient := http.Client{
			Timeout: time.Second * 2,
		}

		req, errReq := http.NewRequest(http.MethodGet, url, nil)
		if errReq != nil {
			fmt.Printf("Requête - Erreur lors de l'initialisation de la requête : %s\n", errReq.Error())
			return []CartesPokemon{}
		}

		res, errRes := httpClient.Do(req)
		if errRes != nil {
			fmt.Printf("Requête - Erreur lors de l'exécution de la requête : %s\n", errRes.Error())
			return []CartesPokemon{}
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusNotFound {
			fmt.Printf("Réponse 404 - Erreur code HTTP : %d, message : %s\n", res.StatusCode, res.Status)
			break
		}

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Réponse - Erreur code HTTP : %d, message : %s\n", res.StatusCode, res.Status)
			return []CartesPokemon{}
		}

		var decodeData CartesPokemon = CartesPokemon{}
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("Decode - Erreur lors du décodage des données : %s\n", errDecode.Error())
			return []CartesPokemon{}
		}
		result = append(result, decodeData)
		j++
	}
	return result
}

// func CardsRequest() []CartesPokemon {
// 	result := []CartesPokemon{}

// 	url := "https://api.tcgdex.net/v2/fr/cards/swsh1-1"

// 	for i := 1; i <= 10; i++ {
// 		j := 1
// 		for {
// 			url = fmt.Sprintf("https://api.tcgdex.net/v2/fr/cards/swsh%d-%d", i, j)

// 			httpClient := http.Client{
// 				Timeout: time.Second * 2,
// 			}

// 			req, errReq := http.NewRequest(http.MethodGet, url, nil)
// 			if errReq != nil {
// 				fmt.Printf("Requête - Erreur lors de l'initialisation de la requête : %s\n", errReq.Error())
// 				return []CartesPokemon{}
// 			}

// 			res, errRes := httpClient.Do(req)
// 			if errRes != nil {
// 				fmt.Printf("Requête - Erreur lors de l'exécution de la requête : %s\n", errRes.Error())
// 				return []CartesPokemon{}
// 			}

// 			defer res.Body.Close()

// 			if res.StatusCode == http.StatusNotFound {
// 				fmt.Printf("Réponse 404 - Erreur code HTTP : %d, message : %s\n", res.StatusCode, res.Status)
// 				break
// 			}

// 			if res.StatusCode != http.StatusOK {
// 				fmt.Printf("Réponse - Erreur code HTTP : %d, message : %s\n", res.StatusCode, res.Status)
// 				return []CartesPokemon{}
// 			}

// 			var decodeData CartesPokemon = CartesPokemon{}
// 			errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
// 			if errDecode != nil {
// 				fmt.Printf("Decode - Erreur lors du décodage des données : %s\n", errDecode.Error())
// 				return []CartesPokemon{}
// 			}
// 			result = append(result, decodeData)
// 			j++
// 		}
// 	}
// 	fmt.Println(result)
// 	return result
// }
