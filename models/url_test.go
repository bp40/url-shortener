package models

import (
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {

	t.Run("Shorten URL function returns correct output", func(t *testing.T) {

		shortenedUrl, err := generateShortUrl()
		if err != nil {
			t.Errorf("Error shortening URL: %s", err)
			return
		}

		if len(shortenedUrl) != 8 {
			t.Errorf("Shortened URL should have 8 characters but got %s which is %d", shortenedUrl, len(shortenedUrl))
		}

	})

}
