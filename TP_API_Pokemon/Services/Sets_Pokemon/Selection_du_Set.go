package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SetPokemon struct {
	CardCount struct {
		Total int `json:"total"`
	} `json:"cardCount"`
	ID          string `json:"id"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
}

func ApiRequest() {

	url := "https://api.tcgdex.net/v2/fr/sets/swsh1"

	for i := 1; i <= 10; i++ {
		url = fmt.Sprintf("https://api.tcgdex.net/v2/en/sets/swsh%d", i)

		httpClient := http.Client{
			Timeout: time.Second * 2,
		}

		req, errReq := http.NewRequest(http.MethodGet, url, nil)
		if errReq != nil {
			fmt.Printf("Requête - Erreur lors de l'initialisation de la requête : %s\n", errReq.Error())
			return
		}

		res, errRes := httpClient.Do(req)
		if errRes != nil {
			fmt.Printf("Requête - Erreur lors de l'exécution de la requête : %s\n", errRes.Error())
			return
		}

		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			fmt.Printf("Réponse - Erreur code HTTP : %d, message : %s\n", res.StatusCode, res.Status)
			return
		}

		var decodeData SetPokemon
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("Decode - Erreur lors du décodage des données : %s\n", errDecode.Error())
			return
		}

		fmt.Printf("Set récupéré :\n")
		fmt.Printf("Nom du Set : %s\n", decodeData.Name)
		fmt.Printf("Série du Set : %s\n", decodeData.ID)
		fmt.Printf("Date de sortie : %s\n", decodeData.ReleaseDate)
		fmt.Printf("Image du Set : %s\n", decodeData.Logo)
		fmt.Printf("Nombre total de cartes : %d\n", decodeData.CardCount.Total)
		fmt.Printf("\n")

	}
}
