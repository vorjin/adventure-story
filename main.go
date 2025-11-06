package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func main() {
	//opening a JSON file
	storyJSONPath := flag.String("json", "gopher.json", "path to the json file to open")
	flag.Parse()

	storyJSON, err := os.ReadFile(*storyJSONPath)
	if err != nil {
		panic(err)
	}

	var storyFile map[string]Chapter
	err = json.Unmarshal([]byte(storyJSON), &storyFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", storyHandler(storyFile))
	if err != nil {
		panic(err)
	}

}

func storyHandler(str map[string]Chapter) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("html/story.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")

		if story, ok := str[path]; ok {
			fmt.Println(story)
			//executing template
			err := tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}

			return
		} else {
			fmt.Printf("Chapter key not found: %s", r.URL.Path)
		}
	}
}
