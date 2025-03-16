package controllers

import (
	"Api/Services/Cartes_Pokemon"
	"Api/templates"
	"Api/utils"
	"net/http"
	"strconv"
)

type PageFavorites struct {
	Cards       []Cartes_Pokemon.CartesPokemon
	TotalPages  int
	CurrentPage int
	NoCardFound bool
}

// Contrôleur pour afficher les cartes en favoris d'un utilisateur avec pagination
func FavorisController(w http.ResponseWriter, r *http.Request) {
	// Simuler un ID utilisateur
	userID := "user123"

	// Récupérer les favoris de l'utilisateur
	favManager := utils.GetFavoritesManager()
	favorites, err := favManager.LoadUserFavorites(userID)
	if err != nil {
		http.Error(w, "Erreur lors du chargement des favoris: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer et gérer la pagination
	page := r.URL.Query().Get("page")
	value, err := strconv.Atoi(page)
	if err != nil || value < 1 {
		value = 1
	}
	cards, totalPages, currentPage := utils.PaginateCards(favorites.Cards, value, 30)

	// Afficher le template avec les données
	data := PageFavorites{
		Cards:       cards,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		NoCardFound: len(favorites.Cards) == 0,
	}
	templates.Temp.ExecuteTemplate(w, "favorites", data)
}

// Contrôleur pour ajouter une carte aux favoris d'un utilisateur
func AddToFavorisController(w http.ResponseWriter, r *http.Request) {
	userID := "user123"

	// Vérifier la méthode et récupérer l'ID de la carte
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	cardID := r.FormValue("card_id")
	if cardID == "" {
		http.Error(w, "ID de carte manquant", http.StatusBadRequest)
		return
	}

	// Récupérer les informations de la carte et l'ajouter aux favoris
	card, err := Cartes_Pokemon.GetCardByID(cardID)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des informations de la carte: "+err.Error(), http.StatusInternalServerError)
		return
	}
	favManager := utils.GetFavoritesManager()
	err = favManager.AddCardToFavorites(userID, card)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout aux favoris: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Rediriger vers la page appropriée
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/favoris"
	}
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

// Contrôleur pour retirer une carte des favoris d'un utilisateur
func RemoveFromFavorisController(w http.ResponseWriter, r *http.Request) {
	userID := "user123"

	// Vérifier la méthode et récupérer l'ID de la carte
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	cardID := r.FormValue("card_id")
	if cardID == "" {
		http.Error(w, "ID de carte manquant", http.StatusBadRequest)
		return
	}

	// Supprimer la carte des favoris
	favManager := utils.GetFavoritesManager()
	err := favManager.RemoveCardFromFavorites(userID, cardID)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression des favoris: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Rediriger vers la page appropriée
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/favoris"
	}
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
