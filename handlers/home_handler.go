package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./static/index.html")

}

func AssetsHandler(writer http.ResponseWriter, request *http.Request) {

	http.ServeFile(writer, request, "./static/assets/"+request.URL.Path[len("/assets/"):])

}
