package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
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

	// parsing the JSON file
	var storyFile map[string]Chapter
	err = json.Unmarshal([]byte(storyJSON), &storyFile)
	if err != nil {
		panic(err)
	}

	//beginning the story
	fmt.Println("Let's begin our adventure!")
	userChoice := storyTeller("intro", storyFile)

	for userChoice != "quit" {
		userChoice = storyTeller(userChoice, storyFile)
	}
}

func storyTeller(chp string, str map[string]Chapter) string {
	if story, ok := str[chp]; ok {
		printStory(story)
		// working with user input
		var index int
		fmt.Scan(&index)

		if index == 0 {
			return "quit"
		} else if index > len(story.Options) {
			fmt.Println("Please, choose the correct option value")
			return chp
		} else {
			trueIndex := index - 1
			userAnswer := story.Options[trueIndex].Arc
			return userAnswer
		}
	} else {
		fmt.Println("Error, the chapter was not found.")
		return "quit"
	}
}

func printStory(story Chapter) {
	for _, text := range story.Story {
		fmt.Println(text)
	}
	fmt.Println()
	// printing the options
	for _, option := range story.Options {
		fmt.Println("Print", slices.Index(story.Options, option)+1, "to:")
		fmt.Println(option.Text)
	}
	fmt.Println("Print 0 to exit program.")
	fmt.Println()
}
