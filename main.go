package main

import (
	"github.com/bp40/url-shortener/handlers"
	"github.com/bp40/url-shortener/middleware"
	"github.com/bp40/url-shortener/models"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sqlx.Open("sqlite3", "./url.db")
	if err != nil {
		log.Fatal(err)
	}

	env := &handlers.Env{
		Urls: models.UrlModel{DB: db},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.HomeHandler)
	mux.HandleFunc("GET /assets/", handlers.AssetsHandler)
	mux.HandleFunc("GET /get/{shortenedURL}", env.UnShortenHandler)
	mux.HandleFunc("POST /shorten", env.ShortenHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	loggedMux := middleware.LoggerMiddleware(mux)

	port := os.Getenv("PORT")

	log.Println("Listening on port " + port)
	err = http.ListenAndServe(":"+port, c.Handler(loggedMux))
	if err != nil {
		log.Fatal(err)
	}

}
