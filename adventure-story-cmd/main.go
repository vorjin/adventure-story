package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
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

	fmt.Println("Let's begin our adventure!")
	storyTeller("intro", storyFile)

	var userChoice string
	fmt.Scan(&userChoice)

	for userChoice != "quit" {
		storyTeller(userChoice, storyFile)
		fmt.Scan(&userChoice)
	}
}

func storyTeller(chp string, str map[string]Chapter) {

	if story, ok := str[chp]; ok {
		for _, text := range story.Story {
			fmt.Println(text)
		}
		for _, option := range story.Options {
			fmt.Println("Print <", option.Arc, "> to do:")
			fmt.Println(option.Text)
		}
		fmt.Println("Print < quit > to exit program.")
		fmt.Println()
	}
}
