package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

const favoritesFilename = "data/favorites.json"

var favoritesMu sync.Mutex

// FavoriteCard représente une carte enregistrée en favori (pour le stockage JSON).
type FavoriteCard struct {
	ID       int    `json:"id"`
	ManaCost int    `json:"manaCost"`
	Attack   int    `json:"attack"`
	Health   int    `json:"health"`
	Rarity   int    `json:"rarity"`
	Name     string `json:"name"`
	Image    string `json:"image"`
}

// RarityName retourne le nom français de la rareté.
func (f FavoriteCard) RarityName() string {
	switch f.Rarity {
	case 1:
		return "Gratuite"
	case 2:
		return "Commune"
	case 3:
		return "Rare"
	case 4:
		return "Épique"
	case 5:
		return "Légendaire"
	default:
		return "Inconnue"
	}
}

// LoadFavorites charge la liste des favoris depuis le fichier JSON.
func LoadFavorites() ([]FavoriteCard, error) {
	favoritesMu.Lock()
	defer favoritesMu.Unlock()

	data, err := os.ReadFile(favoritesFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return []FavoriteCard{}, nil
		}
		return nil, err
	}

	var list []FavoriteCard
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// SaveFavorites enregistre la liste des favoris dans le fichier JSON.
func SaveFavorites(list []FavoriteCard) error {
	favoritesMu.Lock()
	defer favoritesMu.Unlock()

	dir := filepath.Dir(favoritesFilename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(favoritesFilename, data, 0644)
}

// AddFavorite ajoute une carte aux favoris si elle n'y est pas déjà (par ID).
func AddFavorite(card FavoriteCard) error {
	list, err := LoadFavorites()
	if err != nil {
		return err
	}
	for _, c := range list {
		if c.ID == card.ID {
			return nil // déjà en favori
		}
	}
	list = append(list, card)
	return SaveFavorites(list)
}

// RemoveFavorite supprime une carte des favoris par son ID.
func RemoveFavorite(id int) error {
	list, err := LoadFavorites()
	if err != nil {
		return err
	}

	if len(list) == 0 {
		return nil
	}

	newList := make([]FavoriteCard, 0, len(list))
	for _, c := range list {
		if c.ID != id {
			newList = append(newList, c)
		}
	}

	return SaveFavorites(newList)
}
