package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"strings"
)

//go:embed data/actions.txt
var actions string

//go:embed data/characters.txt
var Rcharacters string

//go:embed data/descriptions.txt
var Rdescriptions string

//go:embed data/names.txt
var Rnames string

//go:embed data/settings.txt
var Rsettings string

//go:embed data/plottwists.txt
var Rplottwists string

type markov struct {
	base       string
	next       string
	count      int
	probabilty float64
}

func main() {

	var bookTitle string
	flag.StringVar(&bookTitle, "t", "", "Book title to clean.")
	flag.Parse()

	if bookTitle != "" {
		cleanTitle(bookTitle)
		os.Exit(0)
	}

	characters := strings.Split(Rcharacters, "\n")
	descriptions := strings.Split(Rdescriptions, "\n")
	names := strings.Split(Rnames, "\n")
	settings := strings.Split(Rsettings, "\n")
	actions := strings.Split(actions, "\n")
	plottwists := strings.Split(Rplottwists, "\n")

	dcount := len(descriptions)
	ccount := len(characters)
	ncount := len(names)
	scount := len(settings)
	acount := len(actions)
	pcount := len(plottwists)

	description := descriptions[rand.IntN(dcount)]
	character := characters[rand.IntN(ccount)]
	name := names[rand.IntN(ncount)]
	setting := settings[rand.IntN(scount)]
	action := actions[rand.IntN(acount)]
	plottwist := plottwists[rand.IntN(pcount)]

	save := fmt.Sprintf("|%s %s|%s|%s|%s|%s|\n", character, name, action, description, setting, plottwist)
	// Save to file
	f, err := os.OpenFile("slugs.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(save); err != nil {
		fmt.Println(err)
	}

	fmt.Println(save)

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
