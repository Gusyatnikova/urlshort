package main

import (
	"flag"
	"fmt"
	"github.com/Gusyatnikova/urlshort"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	filePath := flag.String("file", "D:\\GoLang\\urlshort\\main\\redirections.json",
		"an absolute path to json(yaml) file with redirection info.")
	isYAML := flag.Bool("isYaml", false, "set isYaml=True when you want to use .YAML instead of .JSON")
	flag.Parse()

	db := SetupConn()
	defer db.Close()
	gotToRedirect := GetRedirections(db)

	http.HandleFunc("/", home)
	var handler http.HandlerFunc
	mapHandler := urlshort.MapHandler(gotToRedirect, http.DefaultServeMux)
	fileBytes := readFile(*filePath)
	handler = selectHandler(*isYAML, []byte(fileBytes), mapHandler)
	http.ListenAndServe(":3000", http.Handler(handler))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h3>I'm a default serveMux</h3>")
}

func readFile(path string) []byte {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error while read yamlFile: #%v", err)
		os.Exit(1)
	}
	return yamlFile
}

func selectHandler(isYaml bool, fileBytes []byte, mapHandler http.HandlerFunc) http.HandlerFunc {
	var err error
	var handler http.HandlerFunc
	if isYaml {
		handler, err = urlshort.YAMLHandler([]byte(fileBytes), mapHandler)
	} else {
		handler, err = urlshort.JSONHandler([]byte(fileBytes), mapHandler)
	}
	if err != nil {
		panic(err)
	}
	return handler
}
