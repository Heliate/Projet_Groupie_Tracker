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

// Initialiser le gestionnaire de favoris
func InitFavoritesManager(filePath string) {
	favManager = &FavoritesManager{
		filePath: filePath,
	}
}

// Obtenir l'instance du gestionnaire de favoris
func GetFavoritesManager() *FavoritesManager {
	return favManager
}

// Charger les favoris d'un utilisateur depuis le fichier JSON
func (fm *FavoritesManager) LoadUserFavorites(userID string) (UserFavorites, error) {
	fm.mutex.RLock()
	defer fm.mutex.RUnlock()

	// Vérifier si le fichier existe
	if _, err := os.Stat(fm.filePath); os.IsNotExist(err) {
		// Si le fichier n'existe pas, créer un fichier vide avec une structure de base
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

	// Lire le fichier existant
	fileData, err := os.ReadFile(fm.filePath)
	if err != nil {
		return UserFavorites{}, fmt.Errorf("erreur lors de la lecture du fichier de favoris: %v", err)
	}

	// Désérialiser les données
	allFavorites := make(map[string]UserFavorites)
	if err := json.Unmarshal(fileData, &allFavorites); err != nil {
		return UserFavorites{}, fmt.Errorf("erreur lors de la désérialisation des favoris: %v", err)
	}

	// Récupérer les favoris de l'utilisateur, ou créer une entrée vide si elle n'existe pas
	favorites, exists := allFavorites[userID]
	if !exists {
		favorites = UserFavorites{
			UserID: userID,
			Cards:  []Cartes_Pokemon.CartesPokemon{},
		}
		allFavorites[userID] = favorites

		// Sauvegarder la nouvelle structure
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

// Ajouter une carte aux favoris d'un utilisateur
func (fm *FavoritesManager) AddCardToFavorites(userID string, card Cartes_Pokemon.CartesPokemon) error {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()

	// Lire les favoris existants
	fileData, err := os.ReadFile(fm.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Si le fichier n'existe pas, créer un fichier avec la nouvelle carte
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

	// Désérialiser les données
	allFavorites := make(map[string]UserFavorites)
	if err := json.Unmarshal(fileData, &allFavorites); err != nil {
		return fmt.Errorf("erreur lors de la désérialisation des favoris: %v", err)
	}

	// Récupérer les favoris de l'utilisateur, ou créer une entrée vide si elle n'existe pas
	favorites, exists := allFavorites[userID]
	if !exists {
		favorites = UserFavorites{
			UserID: userID,
			Cards:  []Cartes_Pokemon.CartesPokemon{},
		}
	}

	// Vérifier si la carte existe déjà dans les favoris
	for _, existingCard := range favorites.Cards {
		if existingCard.ID == card.ID {
			// La carte est déjà dans les favoris, ne rien faire
			return nil
		}
	}

	// Ajouter la nouvelle carte aux favoris
	favorites.Cards = append(favorites.Cards, card)
	allFavorites[userID] = favorites

	// Sauvegarder la structure mise à jour
	jsonData, err := json.MarshalIndent(allFavorites, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du fichier de favoris: %v", err)
	}

	return os.WriteFile(fm.filePath, jsonData, 0644)
}

// Supprimer une carte des favoris d'un utilisateur
func (fm *FavoritesManager) RemoveCardFromFavorites(userID string, cardID string) error {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()

	// Lire les favoris existants
	fileData, err := os.ReadFile(fm.filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier de favoris: %v", err)
	}

	// Désérialiser les données
	allFavorites := make(map[string]UserFavorites)
	if err := json.Unmarshal(fileData, &allFavorites); err != nil {
		return fmt.Errorf("erreur lors de la désérialisation des favoris: %v", err)
	}

	// Récupérer les favoris de l'utilisateur
	favorites, exists := allFavorites[userID]
	if !exists {
		// L'utilisateur n'a pas de favoris, rien à supprimer
		return nil
	}

	// Créer une nouvelle liste sans la carte à supprimer
	updatedCards := []Cartes_Pokemon.CartesPokemon{}
	for _, card := range favorites.Cards {
		if card.ID != cardID {
			updatedCards = append(updatedCards, card)
		}
	}

	// Mettre à jour la liste des cartes
	favorites.Cards = updatedCards
	allFavorites[userID] = favorites

	// Sauvegarder la structure mise à jour
	jsonData, err := json.MarshalIndent(allFavorites, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du fichier de favoris: %v", err)
	}

	return os.WriteFile(fm.filePath, jsonData, 0644)
}

// Vérifier si une carte est dans les favoris d'un utilisateur
func (fm *FavoritesManager) IsCardInFavorites(userID string, cardID string) (bool, error) {
	fm.mutex.RLock()
	defer fm.mutex.RUnlock()

	favorites, err := fm.LoadUserFavorites(userID)
	if err != nil {
		return false, err
	}

	for _, card := range favorites.Cards {
		if card.ID == cardID {
			return true, nil
		}
	}

	return false, nil
}
