package main

import (
	"fmt"
	"github.com/Gusyatnikova/urlshort"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	gotToRedirect := map[string]string{
		"/google": "https://google.com/",
		"/game":   "https://tetris.com/play-tetris/",
		"/github": "https://github.com/Gusyatnikova/",
	}

	mapHandler := urlshort.MapHandler(gotToRedirect, http.DefaultServeMux)
	yaml := `
- path: /Fun/TV-series
  url: https://disneyplusoriginals.disney.com/show/the-mandalorian
- path: /Fun/The-Child
  url: https://static.wikia.nocookie.net/starwars/images/4/43/TheChild-Fathead.png/revision/latest/scale-to-width-down/500?cb=20201031231040
- path: /Fun/Grogu
  url: https://static.wikia.nocookie.net/starwars/images/4/43/TheChild-Fathead.png/revision/latest/scale-to-width-down/500?cb=20201031231040
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":3000", http.Handler(yamlHandler))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h3>I'm a default serveMux</h3>")
}
