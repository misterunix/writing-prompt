package main

import (
	_ "embed"
	"fmt"
	"math/rand/v2"
	"os"
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

func main() {
	//fmt.Println("Starting the application...")

	characters := strings.Split(Rcharacters, "\n")
	descriptions := strings.Split(Rdescriptions, "\n")
	//names := strings.Split(Rnames, "\n")
	settings := strings.Split(Rsettings, "\n")
	actions := strings.Split(actions, "\n")
	plottwists := strings.Split(Rplottwists, "\n")

	dcount := len(descriptions)
	ccount := len(characters)
	//ncount := len(names)
	scount := len(settings)
	acount := len(actions)
	pcount := len(plottwists)

	description := descriptions[rand.IntN(dcount)]
	character := characters[rand.IntN(ccount)]
	//name := names[rand.IntN(ncount)]
	setting := settings[rand.IntN(scount)]
	action := actions[rand.IntN(acount)]
	plottwist := plottwists[rand.IntN(pcount)]

	//line := description + " " + character + " " + action + " " + plottwist + " " + setting
	// fmt.Println("Generated line:")
	// fmt.Println("Character: ", character)
	// fmt.Println("Action: ", action)
	// fmt.Println("Description: ", description)
	// fmt.Println("Setting: ", setting)
	// fmt.Println("Plottwist: ", plottwist)

	save := fmt.Sprintf("%s, %s, %s, %s, %s\n", character, action, description, setting, plottwist)
	// Save to file
	f, err := os.OpenFile("slugs..txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(save); err != nil {
		fmt.Println(err)
	}

	fmt.Println(save)

}
