package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Gusyatnikova/urlshort"
	_ "github.com/lib/pq" //_ want include without directly reference
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//http://127.0.0.1:49728/?key=9de731ea-c88e-4078-8413-a3e4e9dbf410
const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "Nata2010"
	dbname   = "urlshort"
)

func main() {
	filePath := flag.String("file", "D:\\GoLang\\urlshort\\main\\redirections.json",
		"an absolute path to json(yaml) file with redirection info.")
	isYAML := flag.Bool("isYaml", false, "set isYaml=True when you want to use .YAML instead of .JSON")
	flag.Parse()
	//SQL START HERE
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbPtr, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer dbPtr.Close()
	err = dbPtr.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to Database")
	//table redirection to pathToUrl map[string]string
	rows, err := dbPtr.Query("SELECT path, url FROM urlshort_schema.redirection")
	gotToRedirect := map[string]string{
		"/google": "https://google.com/",
		"/game":   "https://tetris.com/play-tetris/",
	}
	if err != nil {
		log.Println("Cannot SELECT * FROM redirection")
		panic(err)
	}
	sqlMap := make(map[string]string)
	defer rows.Close()
	for rows.Next() {
		var path, url string
		err = rows.Scan(&path, &url)
		if err != nil {
			log.Println("cannot scan rows")
			panic(err)
		}
		sqlMap[path] = url
	}
	err = rows.Err()
	if err == nil {
		gotToRedirect = sqlMap
	}
	fmt.Println(gotToRedirect)


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
