package controllers

import (
	"Api/services/Sets_Pokemon"
	"Api/templates"
	"net/http"
)

func SetPokemonControllers(w http.ResponseWriter, r *http.Request) {

	type DataStruct struct {
		Data []Sets_Pokemon.SetPokemon
	}
	DataToSend := DataStruct{
		Data: Sets_Pokemon.SetRequest(),
	}
	templates.Temp.ExecuteTemplate(w, "sets_pokemon", DataToSend)
}
