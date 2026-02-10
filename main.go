package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"

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

//
//
//

//go:embed data/3/3-a.gz
var R3a []byte

//go:embed data/3/3-c.gz
var R3c []byte

//go:embed data/3/3-d.gz
var R3d []byte

//go:embed data/3/3-n.gz
var R3n []byte

//go:embed data/3/3-p.gz
var R3p []byte

//go:embed data/3/3-s.gz
var R3s []byte

//
//
//

type markov struct {
	base       string
	next       string
	count      int
	probabilty float64
}

var logging bool = false
var mode int

const NUMBER_OF_MODES = 4

var (
	characters   [NUMBER_OF_MODES][]string
	names        [NUMBER_OF_MODES][]string
	actions      [NUMBER_OF_MODES][]string
	descriptions [NUMBER_OF_MODES][]string
	settings     [NUMBER_OF_MODES][]string
	plottwists   [NUMBER_OF_MODES][]string
)

func main() {

	var countWanted int // number of slugs to generate

	modeString := fmt.Sprintf("Mode of operation\n\t0 - Normal (0-%d).\n\t1-Science Fiction\n\t2-Fantasy\n\t3-Paranormal Romance", NUMBER_OF_MODES-1)

	flag.IntVar(&countWanted, "n", 1, "Number of slugs to generate.")
	flag.BoolVar(&logging, "l", false, "Turn logging on.")
	flag.IntVar(&mode, "m", 0, modeString)

	flag.Parse()

	//mode = mode % NUMBER_OF_MODES

	//
	// 0
	//
	characters[0] = PopulateData(R0c, "R0c")
	names[0] = PopulateData(R0n, "R0n")
	actions[0] = PopulateData(R0a, "R0a")
	descriptions[0] = PopulateData(R0d, "R0d")
	settings[0] = PopulateData(R0s, "R0s")
	plottwists[0] = PopulateData(R0p, "R0p")

	characters[1] = PopulateData(R1c, "R1c")
	names[1] = PopulateData(R1n, "R1n")
	actions[1] = PopulateData(R1a, "R1a")
	descriptions[1] = PopulateData(R1d, "R1d")
	settings[1] = PopulateData(R1s, "R1s")
	plottwists[1] = PopulateData(R1p, "R1p")

	characters[2] = PopulateData(R2c, "R2c")
	names[2] = PopulateData(R2n, "R2n")
	actions[2] = PopulateData(R2a, "R2a")
	descriptions[2] = PopulateData(R2d, "R2d")
	settings[2] = PopulateData(R2s, "R2s")
	plottwists[2] = PopulateData(R2p, "R2p")

	characters[3] = PopulateData(R3c, "R3c")
	names[3] = PopulateData(R3n, "R3n")
	actions[3] = PopulateData(R3a, "R3a")
	descriptions[3] = PopulateData(R3d, "R3d")
	settings[3] = PopulateData(R3s, "R3s")
	plottwists[3] = PopulateData(R3p, "R3p")

	//
	//
	//

	// loop to generate slugs
	for range countWanted {

		character, name, action, description, setting, plottwist := createSlug()

		// format slug as CSV line
		save := fmt.Sprintf("Writing Prompt:\nCharacter:'%s'\nName:'%s'\nAction:'%s'\nDescription:'%s'\nSetting:'%s'\nPlot Twist:'%s'\n",
			character, name, action, description, setting, plottwist)

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

func PopulateData(rawdata []byte, name string) []string {

	var final []string

	dcd, err := GzipDecompress(rawdata)
	if err != nil {
		if logging {
			log.Fatalln("Error decompressing", name, ":", err)
		}
		os.Exit(1)
	}
	final = strings.Split(string(dcd), "\n")
	//fmt.Println("Name:", name, "Count:", len(final))
	return final
}
