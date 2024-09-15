package handlers

type Env struct {
	Urls interface {
		GetOriginalUrl(shortenedURL string) (string, error)
		SaveShortURL(originalURL string) (string, error)
	}
}
