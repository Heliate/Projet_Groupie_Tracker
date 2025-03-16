package utils

import (
	"Api/Services/Cartes_Pokemon"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Structure pour stocker les favoris d'un utilisateur
type UserFavorites struct {
	UserID string                         `json:"userId"`
	Cards  []Cartes_Pokemon.CartesPokemon `json:"cards"`
}

// Gestionnaire de favoris avec mutex pour la concurrence
type FavoritesManager struct {
	filePath string
	mutex    sync.RWMutex
}

// Instance unique du gestionnaire de favoris
var favManager *FavoritesManager

// Initialise le gestionnaire de favoris avec le chemin du fichier de stockage
func InitFavoritesManager(filePath string) {
	favManager = &FavoritesManager{
		filePath: filePath,
	}
}

// Retourne l'instance du gestionnaire de favoris
func GetFavoritesManager() *FavoritesManager {
	return favManager
}

// Charge les favoris d'un utilisateur depuis le fichier JSON
func (fm *FavoritesManager) LoadUserFavorites(userID string) (UserFavorites, error) {
	fm.mutex.RLock()
	defer fm.mutex.RUnlock()

	// Crée le fichier s'il n'existe pas
	if _, err := os.Stat(fm.filePath); os.IsNotExist(err) {
		initialData := make(map[string]UserFavorites)
		initialData[userID] = UserFavorites{
			UserID: userID,
			Cards:  []Cartes_Pokemon.CartesPokemon{},
		}

		jsonData, err := json.MarshalIndent(initialData, "", "  ")
		if err != nil {
			return UserFavorites{}, fmt.Errorf("erreur lors de la création du fichier de favoris: %v", err)
		}

		err = os.WriteFile(fm.filePath, jsonData, 0644)
		if err != nil {
			return UserFavorites{}, fmt.Errorf("erreur lors de l'écriture du fichier de favoris: %v", err)
		}

		return initialData[userID], nil
	}

	// Lit le fichier existant
	fileData, err := os.ReadFile(fm.filePath)
	if err != nil {
		return UserFavorites{}, fmt.Errorf("erreur lors de la lecture du fichier de favoris: %v", err)
	}

	// Désérialise les données
	allFavorites := make(map[string]UserFavorites)
	if err := json.Unmarshal(fileData, &allFavorites); err != nil {
		return UserFavorites{}, fmt.Errorf("erreur lors de la désérialisation des favoris: %v", err)
	}

	// Crée une entrée vide si l'utilisateur n'existe pas
	favorites, exists := allFavorites[userID]
	if !exists {
		favorites = UserFavorites{
			UserID: userID,
			Cards:  []Cartes_Pokemon.CartesPokemon{},
		}
		allFavorites[userID] = favorites

		jsonData, err := json.MarshalIndent(allFavorites, "", "  ")
		if err != nil {
			return UserFavorites{}, fmt.Errorf("erreur lors de la mise à jour du fichier de favoris: %v", err)
		}

		err = os.WriteFile(fm.filePath, jsonData, 0644)
		if err != nil {
			return UserFavorites{}, fmt.Errorf("erreur lors de l'écriture du fichier de favoris: %v", err)
		}
	}

	return favorites, nil
}

// Ajoute une carte aux favoris d'un utilisateur
func (fm *FavoritesManager) AddCardToFavorites(userID string, card Cartes_Pokemon.CartesPokemon) error {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()

	// Lit ou crée le fichier de favoris
	fileData, err := os.ReadFile(fm.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			allFavorites := make(map[string]UserFavorites)
			allFavorites[userID] = UserFavorites{
				UserID: userID,
				Cards:  []Cartes_Pokemon.CartesPokemon{card},
			}

			jsonData, err := json.MarshalIndent(allFavorites, "", "  ")
			if err != nil {
				return fmt.Errorf("erreur lors de la création du fichier de favoris: %v", err)
			}

			return os.WriteFile(fm.filePath, jsonData, 0644)
		}
		return fmt.Errorf("erreur lors de la lecture du fichier de favoris: %v", err)
	}

	// Désérialise les données
	allFavorites := make(map[string]UserFavorites)
	if err := json.Unmarshal(fileData, &allFavorites); err != nil {
		return fmt.Errorf("erreur lors de la désérialisation des favoris: %v", err)
	}

	// Récupère ou crée les favoris de l'utilisateur
	favorites, exists := allFavorites[userID]
	if !exists {
		favorites = UserFavorites{
			UserID: userID,
			Cards:  []Cartes_Pokemon.CartesPokemon{},
		}
	}

	// Vérifie si la carte existe déjà
	for _, existingCard := range favorites.Cards {
		if existingCard.ID == card.ID {
			return nil
		}
	}

	// Ajoute la carte aux favoris
	favorites.Cards = append(favorites.Cards, card)
	allFavorites[userID] = favorites

	// Sauvegarde les modifications
	jsonData, err := json.MarshalIndent(allFavorites, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du fichier de favoris: %v", err)
	}

	return os.WriteFile(fm.filePath, jsonData, 0644)
}

// Supprime une carte des favoris d'un utilisateur
func (fm *FavoritesManager) RemoveCardFromFavorites(userID string, cardID string) error {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()

	// Lit le fichier de favoris
	fileData, err := os.ReadFile(fm.filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier de favoris: %v", err)
	}

	// Désérialise les données
	allFavorites := make(map[string]UserFavorites)
	if err := json.Unmarshal(fileData, &allFavorites); err != nil {
		return fmt.Errorf("erreur lors de la désérialisation des favoris: %v", err)
	}

	// Vérifie si l'utilisateur a des favoris
	favorites, exists := allFavorites[userID]
	if !exists {
		return nil
	}

	// Filtre la liste pour retirer la carte spécifiée
	updatedCards := []Cartes_Pokemon.CartesPokemon{}
	for _, card := range favorites.Cards {
		if card.ID != cardID {
			updatedCards = append(updatedCards, card)
		}
	}

	// Met à jour la liste des cartes
	favorites.Cards = updatedCards
	allFavorites[userID] = favorites

	// Sauvegarde les modifications
	jsonData, err := json.MarshalIndent(allFavorites, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du fichier de favoris: %v", err)
	}

	return os.WriteFile(fm.filePath, jsonData, 0644)
}

// Vérifie si une carte est dans les favoris d'un utilisateur
func (fm *FavoritesManager) IsCardInFavorites(userID string, cardID string) (bool, error) {
	fm.mutex.RLock()
	defer fm.mutex.RUnlock()

	// Charge les favoris de l'utilisateur
	favorites, err := fm.LoadUserFavorites(userID)
	if err != nil {
		return false, err
	}

	// Recherche la carte dans les favoris
	for _, card := range favorites.Cards {
		if card.ID == cardID {
			return true, nil
		}
	}

	return false, nil
}
