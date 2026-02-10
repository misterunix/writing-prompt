package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

//go:embed data/0/characters.txt
var Rcharacters string

//go:embed data/0/actions.txt
var Ractions string

//go:embed data/0/descriptions.txt
var Rdescriptions string

//go:embed data/0/names.txt
var Rnames string

//go:embed data/0/settings.txt
var Rsettings string

//go:embed data/0/plottwists.txt
var Rplottwists string

//
//
//

//go:embed data/1/actions.txt
var RscifiActions string

//go:embed data/1/characters.txt
var RscifiCharacters string

//go:embed data/1/descriptions.txt
var RscifiDescriptions string

//go:embed data/1/names.txt
var RscifiNames string

//go:embed data/1/settings.txt
var RscifiSettings string

//go:embed data/1/plottwists.txt
var RscifiPlottwists string

//var workingDir string

type markov struct {
	base       string
	next       string
	count      int
	probabilty float64
}

var port int = 5544
var logging bool = false

var (
	characters   []string
	names        []string
	actions      []string
	descriptions []string
	settings     []string
	plottwists   []string
)

func main() {

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var countWanted int // number of slugs to generate
	var bookTitle string
	var mode int

	// parse command line arguments
	flag.IntVar(&countWanted, "n", 1, "Number of slugs to generate.")
	flag.StringVar(&bookTitle, "t", "", "Book title to clean.")
	flag.BoolVar(&logging, "l", false, "Turn logging on.")
	flag.IntVar(&port, "p", 5544, "Port number for web server.")
	flag.IntVar(&mode, "m", 0, "Mode of operation: 0 for default, 1 for sci-fi.")

	flag.Parse()

	// WTF did I do?
	if bookTitle != "" {
		cleanTitle(bookTitle)
		os.Exit(0)
	}

	dataDirBase := path.Join(workingDir, "data")
	dataFolder := strconv.Itoa(mode)
	dataDir := path.Join(dataDirBase, dataFolder)

	// check if data directory exists
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		fmt.Printf("Data directory '%s' does not exist.\n", dataDir)
		os.Exit(1)
	}

	actionsRaw, err := os.ReadFile(path.Join(dataDir, "actions.txt"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	actions = strings.Split(string(actionsRaw), "\n")
	charactersRaw, err := os.ReadFile(path.Join(dataDir, "characters.txt"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	characters = strings.Split(string(charactersRaw), "\n")
	descriptionsRaw, err := os.ReadFile(path.Join(dataDir, "descriptions.txt"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	descriptions = strings.Split(string(descriptionsRaw), "\n")
	namesRaw, err := os.ReadFile(path.Join(dataDir, "names.txt"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	names = strings.Split(string(namesRaw), "\n")
	plottwistsRaw, err := os.ReadFile(path.Join(dataDir, "plottwists.txt"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plottwists = strings.Split(string(plottwistsRaw), "\n")
	settingsRaw, err := os.ReadFile(path.Join(dataDir, "settings.txt"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	settings = strings.Split(string(settingsRaw), "\n")

	// get home directory
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// create data directory
	csvDir := path.Join(workingDir, "writing-prompts")
	err = os.MkdirAll(csvDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Data directory: '", csvDir, "'\n")

	// create temp file in data directory
	f, err := os.CreateTemp(csvDir, "slugs-*.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// loop to generate slugs
	for range countWanted {

		character, name, action, description, setting, plottwist := createSlug()

		/*
			description := descriptions[rand.IntN(dcount)]
			character := characters[rand.IntN(ccount)]
			name := names[rand.IntN(ncount)]
			setting := settings[rand.IntN(scount)]
			action := actions[rand.IntN(acount)]
			plottwist := plottwists[rand.IntN(pcount)]
		*/
		// format slug as CSV line
		save := fmt.Sprintf("Writing Prompt:\nCharacter:'%s'\nName:'%s'\nAction:'%s'\nDescription:'%s'\nSetting:'%s'\nPlot Twist:'%s'\n",
			character, name, action, description, setting, plottwist)

		// write slug to file
		if _, err := f.WriteString(save); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// print slug to console
		fmt.Printf("%s", save)
	}
}

func createSlug() (string, string, string, string, string, string) {
	dcount := len(descriptions)
	ccount := len(characters)
	ncount := len(names)
	scount := len(settings)
	acount := len(actions)
	pcount := len(plottwists)

	character := characters[rand.IntN(ccount)]
	name := names[rand.IntN(ncount)]
	setting := settings[rand.IntN(scount)]
	action := actions[rand.IntN(acount)]
	description := descriptions[rand.IntN(dcount)]
	plottwist := plottwists[rand.IntN(pcount)]
	return character, name, action, description, setting, plottwist
}

func cleanTitle(filename string) error {

	var titles []string // tempory slice to hold titles

	f, err := os.ReadFile(filename) // binary read file
	if err != nil {
		fmt.Println(err)
		return err
	}

	lines := strings.Split(string(f), "\n") // split file into lines

	// loop through lines and strip white space
	for _, line := range lines {

		striped := strings.TrimSpace(line)

		// skip empty lines
		if striped == "" {
			continue
		}

		// skip single word titles
		if len(striped) == 1 {
			continue
		}

		paren := false

		for {

			if strings.Contains(striped, "(") {
				p1 := strings.Index(striped, "(")
				p2 := strings.Index(striped, ")")
				if p1 == -1 || p2 == -1 {
					break
				}
				striped = striped[:p1-1] + striped[p2+1:]
				paren = true
			} else {
				paren = false
			}
			if !paren {
				break
			}
		}

		// split title and author
		t := strings.Split(striped, " by ")

		// add title to slice
		titles = append(titles, t[0])
	}

	sort.StringSlice(titles).Sort() // sort titles

	l := len(titles) // get length of titles

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if titles[i] == titles[j] {
				titles = append(titles[:j], titles[j+1:]...)
				j--
				l--
			}
		}
	}

	chain := markovChain(titles) // create markov chain
	for k, v := range chain {
		fmt.Println(k, v)
	}

	// save to file
	os.WriteFile("cleaned.txt", []byte(strings.Join(titles, "\n")), 0644)
	return nil
}

// Create Markov chain from a slice of strings
func markovChain(data []string) map[string][]string {
	chain := make(map[string][]string)

	for _, line := range data {
		words := strings.Split(line, " ")
		for i := 0; i < len(words)-1; i++ {
			key := words[i]
			if _, ok := chain[key]; !ok {
				chain[key] = []string{}
			}
			chain[key] = append(chain[key], words[i+1])
		}
	}
	return chain
}

// // Create Markov chain from a slice of strings
// func markovChain2(data []string) map[string][]markov {
// 	chain := make(map[string][]string)

// 	for _, line := range data {
// 		words := strings.Split(line, " ")
// 		for i := 0; i < len(words)-1; i++ {
// 			key := words[i]
// 			if _, ok := chain[key]; !ok {
// 				chain[key] = []string{}
// 			}
// 			chain[key] = append(chain[key], words[i+1])
// 		}
// 	}
// 	return chain
// }
