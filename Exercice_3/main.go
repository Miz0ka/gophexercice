package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "html/template"
	"log"
	"net/http"
	"os"
	"text/template"
	_ "time"
)

type StoryTelling struct {
	Title   string     `json:"title"`
	Story   []string   `json:"story"`
	Options []OptionsT `json:"options"`
}

type OptionsT struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Storys map[string]StoryTelling

func main() {
	var storys Storys

	fmt.Println("Parsing Json start...")
	fmt.Println(" - open file")
	f, e := os.Open("./gopher.json")
	if e != nil {
		panic(e)
	}
	fmt.Println(" - Decode file")
	d := json.NewDecoder(f)
	if err := d.Decode(&storys); err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", storys)

	mux := http.NewServeMux()

	StoryHandler := JsonHandler(storys, mux)

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", StoryHandler))
}

func JsonHandler(pathsToUrls Storys, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		fmt.Println("Searching template : ", path)
		if dest, ok := pathsToUrls[path]; ok {
			fmt.Println("Loading template : ", path)
			templates := template.Must(template.ParseFiles("./templates/template.html"))
			err := templates.Execute(w, dest)

			if err != nil {
				log.Fatalf("Template execution: %s", err)
			}
			return
		} else {
			http.Redirect(w, r, "/intro", http.StatusPermanentRedirect)
			return
		}
	}
}
