package main

import (
	"flag"
	"fmt"
	"github.com/Gusyatnikova/urlshort"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	yamlFilePath := flag.String("yaml", "C:\\redirections.yaml",
		"an absolute path to yaml file with redirection info. Example:\n" +
			"- path: <your-path>\nurl: <url-to-redirect>\n")
	flag.Parse()
	http.HandleFunc("/", home)
	gotToRedirect := map[string]string{
		"/google": "https://google.com/",
		"/game":   "https://tetris.com/play-tetris/",
		"/github": "https://github.com/Gusyatnikova/",
	}
	mapHandler := urlshort.MapHandler(gotToRedirect, http.DefaultServeMux)
	yaml := readYamlFile(*yamlFilePath)
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

func readYamlFile(path string) []byte {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error while read yamlFile: #%v", err)
		os.Exit(1)
	}
	return yamlFile
}
