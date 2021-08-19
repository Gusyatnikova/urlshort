package main

import (
	"fmt"
	"github.com/Gusyatnikova/urlshort"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	gotToRedirect := map[string]string {
		"/google": "https://google.com/",
		"/game": "https://tetris.com/play-tetris/",
		"/github": "https://github.com/Gusyatnikova/",
	}

	http.ListenAndServe(
		":3000",
		urlshort.MapHandler(gotToRedirect, http.DefaultServeMux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h3>I'm a default serveMux</h3>")
}