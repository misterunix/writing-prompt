package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"sort"
	"strings"
)

//go:embed data/0/0-a.gz
var R0a []byte

//go:embed data/0/0-c.gz
var R0c []byte

//go:embed data/0/0-d.gz
var R0d []byte

//go:embed data/0/0-n.gz
var R0n []byte

//go:embed data/0/0-p.gz
var R0p []byte

//go:embed data/0/0-s.gz
var R0s []byte

//
//
//

//go:embed data/1/1-a.gz
var R1a []byte

//go:embed data/1/1-c.gz
var R1c []byte

//go:embed data/1/1-d.gz
var R1d []byte

//go:embed data/1/1-n.gz
var R1n []byte

//go:embed data/1/1-p.gz
var R1p []byte

//go:embed data/1/1-s.gz
var R1s []byte

//
//
//

//go:embed data/2/2-a.gz
var R2a []byte

//go:embed data/2/2-c.gz
var R2c []byte

//go:embed data/2/2-d.gz
var R2d []byte

//go:embed data/2/2-n.gz
var R2n []byte

//go:embed data/2/2-p.gz
var R2p []byte

//go:embed data/2/2-s.gz
var R2s []byte

//var workingDir string

type markov struct {
	base       string
	next       string
	count      int
	probabilty float64
}

var port int = 5544
var logging bool = false
var mode int

var (
	characters   [3][]string
	names        [3][]string
	actions      [3][]string
	descriptions [3][]string
	settings     [3][]string
	plottwists   [3][]string
)

func main() {

	// workingDir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	var countWanted int // number of slugs to generate
	var bookTitle string

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

	//
	// 0
	//
	dcd, err := GzipDecompress(R0c)
	if err != nil {
		fmt.Println("Error decompressing R0c:", err)
		os.Exit(1)
	}
	characters[0] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R0n)
	if err != nil {
		fmt.Println("Error decompressing R0n:", err)
		os.Exit(1)
	}
	names[0] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R0a)
	if err != nil {
		fmt.Println("Error decompressing R0a:", err)
		os.Exit(1)
	}
	actions[0] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R0d)
	if err != nil {
		fmt.Println("Error decompressing R0d:", err)
		os.Exit(1)
	}
	descriptions[0] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R0s)
	if err != nil {
		fmt.Println("Error decompressing R0s:", err)
		os.Exit(1)
	}
	settings[0] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R0p)
	if err != nil {
		fmt.Println("Error decompressing R0p:", err)
		os.Exit(1)
	}
	plottwists[0] = strings.Split(string(dcd), "\n")

	//
	//
	//

	dcd, err = GzipDecompress(R1c)
	if err != nil {
		fmt.Println("Error decompressing R1c:", err)
		os.Exit(1)
	}
	characters[1] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R1n)
	if err != nil {
		fmt.Println("Error decompressing R1n:", err)
		os.Exit(1)
	}
	names[1] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R1a)
	if err != nil {
		fmt.Println("Error decompressing R1a:", err)
		os.Exit(1)
	}
	actions[1] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R1d)
	if err != nil {
		fmt.Println("Error decompressing R1d:", err)
		os.Exit(1)
	}
	descriptions[1] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R1s)
	if err != nil {
		fmt.Println("Error decompressing R1s:", err)
		os.Exit(1)
	}
	settings[1] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R1p)
	if err != nil {
		fmt.Println("Error decompressing R1p:", err)
		os.Exit(1)
	}
	plottwists[1] = strings.Split(string(dcd), "\n")

	//
	//
	//

	dcd, err = GzipDecompress(R2c)
	if err != nil {
		fmt.Println("Error decompressing R2c:", err)
		os.Exit(1)
	}
	characters[2] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R2n)
	if err != nil {
		fmt.Println("Error decompressing R2n:", err)
		os.Exit(1)
	}
	names[2] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R2a)
	if err != nil {
		fmt.Println("Error decompressing R2a:", err)
		os.Exit(1)
	}
	actions[2] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R2d)
	if err != nil {
		fmt.Println("Error decompressing R2d:", err)
		os.Exit(1)
	}
	descriptions[2] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R2s)
	if err != nil {
		fmt.Println("Error decompressing R2s:", err)
		os.Exit(1)
	}
	settings[2] = strings.Split(string(dcd), "\n")

	dcd, err = GzipDecompress(R2p)
	if err != nil {
		fmt.Println("Error decompressing R2p:", err)
		os.Exit(1)
	}
	plottwists[2] = strings.Split(string(dcd), "\n")

	//
	//
	//

	// create data directory
	// csvDir := path.Join(workingDir, "writing-prompts")
	// err = os.MkdirAll(csvDir, os.ModePerm)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Print("Data directory: '", csvDir, "'\n")

	// // create temp file in data directory
	// f, err := os.CreateTemp(csvDir, "slugs-*.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()

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
		// if _, err := f.WriteString(save); err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		//}

		// print slug to console
		fmt.Printf("%s", save)
	}
}

func createSlug() (string, string, string, string, string, string) {
	dcount := len(descriptions[mode])
	ccount := len(characters[mode])
	ncount := len(names[mode])
	scount := len(settings[mode])
	acount := len(actions[mode])
	pcount := len(plottwists[mode])

	character := characters[mode][rand.IntN(ccount)]
	name := names[mode][rand.IntN(ncount)]
	setting := settings[mode][rand.IntN(scount)]
	action := actions[mode][rand.IntN(acount)]
	description := descriptions[mode][rand.IntN(dcount)]
	plottwist := plottwists[mode][rand.IntN(pcount)]
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

// GzipDecompress decompresses a gzip-compressed byte slice.
func GzipDecompress(data []byte) ([]byte, error) {
	// Wrap the byte slice in a bytes.Reader to use it as an io.Reader
	buf := bytes.NewReader(data)

	// Create a new gzip reader
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer reader.Close()

	// Read all the decompressed data into a new byte slice
	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read decompressed data: %w", err)
	}

	return decompressed, nil
}
