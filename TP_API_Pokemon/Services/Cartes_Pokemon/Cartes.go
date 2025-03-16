package Cartes_Pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CartesPokemon struct {
	ID         string   `json:"id"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Rarity     string   `json:"rarity"`
	Hp         int      `json:"hp"`
	Types      []string `json:"types"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
}

func CardsBySetRequest(name string) []CartesPokemon {
	result := []CartesPokemon{}
	j := 1
	for {
		url := fmt.Sprintf("https://api.tcgdex.net/v2/fr/cards/%s-%d", name, j)

		httpClient := http.Client{
			Timeout: time.Second * 20,
		}

		req, errReq := http.NewRequest(http.MethodGet, url, nil)
		if errReq != nil {
			fmt.Printf("Requête - Erreur lors de l'initialisation de la requête : %s\n", errReq.Error())
			return []CartesPokemon{}
		}

		req.Header.Set("User-Agent", "Ynov Campus module groupie tracker")

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

// GetCardByID récupère une carte spécifique par son ID
func GetCardByID(cardID string) (CartesPokemon, error) {
	url := fmt.Sprintf("https://api.tcgdex.net/v2/fr/cards/%s", cardID)

	httpClient := http.Client{
		Timeout: time.Second * 20,
	}

	req, errReq := http.NewRequest(http.MethodGet, url, nil)
	if errReq != nil {
		return CartesPokemon{}, fmt.Errorf("erreur lors de l'initialisation de la requête : %s", errReq.Error())
	}

	req.Header.Set("User-Agent", "Ynov Campus module groupie tracker")

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		return CartesPokemon{}, fmt.Errorf("erreur lors de l'exécution de la requête : %s", errRes.Error())
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return CartesPokemon{}, fmt.Errorf("carte non trouvée (ID: %s)", cardID)
	}

	if res.StatusCode != http.StatusOK {
		return CartesPokemon{}, fmt.Errorf("erreur code HTTP : %d, message : %s", res.StatusCode, res.Status)
	}

	var card CartesPokemon
	errDecode := json.NewDecoder(res.Body).Decode(&card)
	if errDecode != nil {
		return CartesPokemon{}, fmt.Errorf("erreur lors du décodage des données : %s", errDecode.Error())
	}

	return card, nil
}
