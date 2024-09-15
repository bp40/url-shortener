package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

func decodeAndVerifyOriginalURL(r *http.Request) (*url.URL, error) {

	decoder := json.NewDecoder(r.Body)
	var rawUrl struct {
		Url string `json:"original_url"`
	}
	err := decoder.Decode(&rawUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rawUrl.Url == "" {
		return nil, errors.New("URL is required")
	}

	u, err := url.ParseRequestURI(rawUrl.Url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !urlIsReachable(u) {
		return nil, errors.New("url is not reachable")
	}

	return u, nil

}

func urlIsReachable(url *url.URL) bool {

	_, err := http.Get(url.String())
	if err != nil {
		return false
	}

	return true
}

func (env *Env) ShortenHandler(w http.ResponseWriter, r *http.Request) {

	rawUrl, err := decodeAndVerifyOriginalURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	if rawUrl == nil {
		http.Error(w, "Error parsing URL", http.StatusInternalServerError)
		return
	}

	shortenedURL, err := env.Urls.SaveShortURL(rawUrl.String())

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"url":"` + shortenedURL + `"}`))
}

func (env *Env) UnShortenHandler(w http.ResponseWriter, r *http.Request) {
	originalURL, err := env.Urls.GetOriginalUrl(r.PathValue("shortenedURL"))
	if err != nil {
		log.Printf(err.Error())
		w.Write([]byte(`{"original_url": ""}`))
	}
	w.Write([]byte(`{"original_url":"` + originalURL + `"}`))
}
