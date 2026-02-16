package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Token est le token d'authentification pour l'API Blizzard
var Token string = "EUphcku68aRHkfRoGFd1GE3i35EbCZDIbf"

type LocalizedString struct {
	DE string `json:"de_DE"`
	US string `json:"en_US"`
	ES string `json:"es_ES"`
	MX string `json:"es_MX"`
	FR string `json:"fr_FR"`
	IT string `json:"it_IT"`
	JP string `json:"ja_JP"`
	KR string `json:"ko_KR"`
	PL string `json:"pl_PL"`
	BR string `json:"pt_BR"`
	RU string `json:"ru_RU"`
	TH string `json:"th_TH"`
	CN string `json:"zh_CN"`
	TW string `json:"zh_TW"`
}

// FlexibleLocalizedString peut désérialiser soit une string soit un objet LocalizedString
// Utilise un champ anonyme (embedded) pour permettre l'accès direct aux champs (ex: .FR)
type FlexibleLocalizedString struct {
	LocalizedString
}

// UnmarshalJSON permet de gérer les deux formats : string ou objet
func (f *FlexibleLocalizedString) UnmarshalJSON(data []byte) error {
	// Essayer d'abord comme une string
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		// Si c'est une string, on la met dans le champ FR
		f.FR = str
		return nil
	}
	// Sinon, essayer comme un objet LocalizedString
	return json.Unmarshal(data, &f.LocalizedString)
}

type Card struct {
	ID       int                     `json:"id"`
	ManaCost int                     `json:"manaCost"`
	Attack   int                     `json:"attack"`
	Health   int                     `json:"health"`
	Rarity   int                     `json:"rarityId"`
	Name     FlexibleLocalizedString `json:"name"`
	Image    FlexibleLocalizedString `json:"image"`
}

// RarityName retourne le nom français de la rareté
func (c Card) RarityName() string {
	switch c.Rarity {
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

type AllCards struct {
	CardCount int    `json:"cardCount"`
	PageCount int    `json:"pageCount"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
	Cards     []Card `json:"cards"`
}

// Deck représente un deck Hearthstone.
type Deck struct {
	Format string `json:"format"`
	Cards  []Card `json:"cards"`
}

// CardBack représente un dos de carte Hearthstone.
type CardBack struct {
	ID    int             `json:"id"`
	Name  LocalizedString `json:"name"`
	Image string          `json:"image"`
}

// AllCardBacks représente la réponse paginée de l'API pour les dos de cartes.
type AllCardBacks struct {
	CardbackCount int        `json:"cardbackCount"`
	PageCount     int        `json:"pageCount"`
	Page          int        `json:"page"`
	PageSize      int        `json:"pageSize"`
	CardBacks     []CardBack `json:"cardBacks"`
}

// GetCardsPage récupère une page de cartes directement depuis l'API Hearthstone.
// Le paramètre textFilter permet de filtrer les cartes sur leur texte (effet, nom, etc.).
func GetCardsPage(page int, pageSize int, textFilter string) (*AllCards, int, error) {
	_client := http.Client{
		Timeout: 20 * time.Second,
	}

	baseURL := "https://eu.api.blizzard.com/hearthstone/cards"

	params := url.Values{}
	params.Set("page", strconv.Itoa(page))
	params.Set("pageSize", strconv.Itoa(pageSize))
	params.Set("locale", "fr_FR")
	if textFilter != "" {
		params.Set("textFilter", textFilter)
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, reqErr := http.NewRequest(http.MethodGet, fullURL, nil)
	if reqErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetCardsPage - Erreur lors de la préparation de la requête : %s", reqErr)
	}

	req.Header.Set("Authorization", "Bearer "+Token)

	res, resErr := _client.Do(req)
	if resErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetCardsPage - Erreur lors de l'envoi de la requête : %s", resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// Lire le corps de la réponse pour obtenir plus de détails sur l'erreur
		errorBody, _ := io.ReadAll(res.Body)
		return nil, res.StatusCode, fmt.Errorf("GetCardsPage - Erreur dans la réponse code : %d, message : %s, body : %s", res.StatusCode, res.Status, string(errorBody))
	}

	var data AllCards

	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetCardsPage - Erreur lors du décodage des données : %s", decodeErr.Error())
	}

	return &data, http.StatusOK, nil
}

// GetCardBacksPage récupère une page de dos de cartes directement depuis l'API Hearthstone.
func GetCardBacksPage(page int, pageSize int) (*AllCardBacks, int, error) {
	_client := http.Client{
		Timeout: 20 * time.Second,
	}

	baseURL := "https://eu.api.blizzard.com/hearthstone/cardbacks"

	url := fmt.Sprintf("%s?page=%d&pageSize=%d", baseURL, page, pageSize)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetCardBacksPage - Erreur lors de la préparation de la requête : %s", reqErr)
	}

	req.Header.Set("Authorization", "Bearer "+Token)

	res, resErr := _client.Do(req)
	if resErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetCardBacksPage - Erreur lors de l'envoi de la requête : %s", resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("GetCardBacksPage - Erreur dans la réponse code : %d, message : %s", res.StatusCode, res.Status)
	}

	var data AllCardBacks

	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetCardBacksPage - Erreur lors du décodage des données : %s", decodeErr.Error())
	}

	return &data, http.StatusOK, nil
}

// GetDeckPage récupère les informations d'un deck Hearthstone à partir de son code
// en utilisant l'endpoint de l'API Hearthstone :
// https://eu.api.blizzard.com/hearthstone/deck?code=<deckCode>
func GetDeckPage(deckCode string) (*Deck, int, error) {
	_client := http.Client{
		Timeout: 20 * time.Second,
	}

	baseURL := "https://eu.api.blizzard.com/hearthstone/deck"

	URLDeckCode := url.QueryEscape(deckCode)
	fullURL := fmt.Sprintf("%s?code=%s", baseURL, URLDeckCode)

	req, reqErr := http.NewRequest(http.MethodGet, fullURL, nil)
	if reqErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetDeckPage - Erreur lors de la préparation de la requête : %s", reqErr)
	}

	req.Header.Set("Authorization", "Bearer "+Token)

	res, resErr := _client.Do(req)
	if resErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetDeckPage - Erreur lors de l'envoi de la requête : %s", resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("GetDeckPage - Erreur dans la réponse code : %d, message : %s", res.StatusCode, res.Status)
	}

	var deck Deck

	decodeErr := json.NewDecoder(res.Body).Decode(&deck)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetDeckPage - Erreur lors du décodage des données : %s", decodeErr.Error())
	}

	return &deck, http.StatusOK, nil
}
