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

// Contrôleur pour afficher les favoris d'un utilisateur
func FavorisController(w http.ResponseWriter, r *http.Request) {
	// Dans un environnement réel, vous récupéreriez l'ID utilisateur depuis la session
	// Pour cet exemple, nous utilisons un ID fixe
	userID := "user123"

	// Récupérer le gestionnaire de favoris
	favManager := utils.GetFavoritesManager()

	// Récupérer les favoris de l'utilisateur
	favorites, err := favManager.LoadUserFavorites(userID)
	if err != nil {
		http.Error(w, "Erreur lors du chargement des favoris: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer la page demandée
	page := r.URL.Query().Get("page")
	value, err := strconv.Atoi(page)
	if err != nil || value < 1 {
		value = 1
	}

	// Paginer les cartes favorites
	cards, totalPages, currentPage := utils.PaginateCards(favorites.Cards, value, 30)

	// Préparer les données pour le template
	data := PageFavorites{
		Cards:       cards,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		NoCardFound: len(favorites.Cards) == 0,
	}

	// Afficher le template des favoris
	templates.Temp.ExecuteTemplate(w, "favorites", data)
}

// Contrôleur pour ajouter une carte aux favoris
func AddToFavorisController(w http.ResponseWriter, r *http.Request) {
	// Dans un environnement réel, vous récupéreriez l'ID utilisateur depuis la session
	userID := "user123"

	// Vérifier que la méthode est POST
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer l'ID de la carte depuis le formulaire
	cardID := r.FormValue("card_id")
	if cardID == "" {
		http.Error(w, "ID de carte manquant", http.StatusBadRequest)
		return
	}

	// Récupérer les informations de la carte
	card, err := Cartes_Pokemon.GetCardByID(cardID)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des informations de la carte: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer le gestionnaire de favoris
	favManager := utils.GetFavoritesManager()

	// Ajouter la carte aux favoris
	err = favManager.AddCardToFavorites(userID, card)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout aux favoris: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Rediriger vers la page d'origine (utiliser un paramètre de redirection si nécessaire)
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/favoris"
	}

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

// Contrôleur pour supprimer une carte des favoris
func RemoveFromFavorisController(w http.ResponseWriter, r *http.Request) {
	// Dans un environnement réel, vous récupéreriez l'ID utilisateur depuis la session
	userID := "user123"

	// Vérifier que la méthode est POST
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer l'ID de la carte depuis le formulaire
	cardID := r.FormValue("card_id")
	if cardID == "" {
		http.Error(w, "ID de carte manquant", http.StatusBadRequest)
		return
	}

	// Récupérer le gestionnaire de favoris
	favManager := utils.GetFavoritesManager()

	// Supprimer la carte des favoris
	err := favManager.RemoveCardFromFavorites(userID, cardID)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression des favoris: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Rediriger vers la page d'origine (utiliser un paramètre de redirection si nécessaire)
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/favoris"
	}

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
