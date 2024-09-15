package models

import (
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
)

type UrlModel struct {
	DB *sqlx.DB
}

type Url struct {
	Id          uint      `json:"id" db:"id"`
	ShortUrl    string    `json:"short_url" db:"short_url"`
	OriginalUrl string    `json:"original_url" db:"original_url"`
	CreatedAt   time.Time `json:"creation_date" db:"creation_date"`
	ExpiresAt   time.Time `json:"expiration_date" db:"expiration_date"`
}

func (m UrlModel) GetOriginalUrl(shortenedURL string) (string, error) {

	url := Url{}

	err := m.DB.Get(&url, "SELECT original_url FROM Urls WHERE short_url = ?", shortenedURL)
	if err != nil {
		return "", err
	}

	return url.OriginalUrl, nil

}

func (m UrlModel) SaveShortURL(originalURL string) (string, error) {

	shortURL, err := generateShortUrl()
	if err != nil {
		return "", err
	}

	query := `INSERT INTO urls (short_url, original_url, creation_date, expiration_date) VALUES (?, ?, ?, ?)`

	m.DB.MustExec(query, shortURL, originalURL, time.Now(), time.Now().Add(time.Hour*24))

	return shortURL, nil
}

func generateShortUrl() (string, error) {

	urlLength := 8
	charset := []rune("23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ") // BASE 56

	shortURL := make([]rune, urlLength)
	for i := range urlLength {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortURL), nil
}
