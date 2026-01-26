package main

/*
import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootPage)
	mux.HandleFunc("/css/", cssPage)
	mux.HandleFunc("/js/", jsPage)
	mux.HandleFunc("/images/", imagePage)

	//mux.HandleFunc("POST /roll/", rollbones)

	hostport := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	http.ListenAndServe(hostport, mux)
}

func rootPage(w http.ResponseWriter, r *http.Request) {

	pathStr := path.Base(r.URL.Path)

	if logging {
		fmt.Println("Request for:", pathStr)
	}

	type Slug struct {
		Character   string
		Name        string
		Action      string
		Description string
		Setting     string
		Plottwist   string
	}

var Slugs []Slug

character, name, action, description, setting, plottwist := createSlug()
save := fmt.Sprintf("Writing Prompt:\nCharacter:'%s'\nName:'%s'\nAction:'%s'\nDescription:'%s'\nSetting:'%s'\nPlot Twist:'%s'\n",
			character, name, action, description, setting, plottwist)


	w.Write([]byte(indexHTML))

}

func cssPage(w http.ResponseWriter, r *http.Request) {
	pathStr := path.Base(r.URL.Path)

	if logging {
		fmt.Println("Request for:", pathStr)
	}

	if strings.Contains(pathStr, "w3.css") {
		w.Header().Set("Content-Type", "text/css")
		w.Header().Set("Cache-Control", "no-cache")
		w.Write([]byte(w3CSS))
		return
	}

	//w.Header().Set("Content-Type", "text/plain")

	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "no-cache")

	contents, err := os.ReadFile(path.Join("webfiles", pathStr))
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Write(contents)

}

func imagePage(w http.ResponseWriter, r *http.Request) {
	path := path.Base(r.URL.Path)

	if logging {
		fmt.Println("Request for:", path)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-cache")

	contents, err := os.ReadFile("webfiles/" + path)
	if err != nil {

		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Write(contents)
}

func jsPage(w http.ResponseWriter, r *http.Request) {
	pathStr := path.Base(r.URL.Path)

	if logging {
		fmt.Println("js Request for:", pathStr)
	}

	if strings.Contains(pathStr, "htmx.min.js") {
		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("Cache-Control", "no-cache")
		w.Write([]byte(htmxJS))
		return
	}

}

func rollbones(w http.ResponseWriter, r *http.Request) {

	pathStr := path.Base(r.URL.Path)

	if logging {
		fmt.Println("roll Request for:", pathStr)
	}

	p := r.FormValue("diceroll")

	terminate := IsExitCommand(p)
	if terminate {
		os.Exit(0)
	}

	if logging {
		fmt.Println("Dice Roll Pattern:", p)
	}

	numDice, dieSize, modifier, top, err := ParseDiceNotation(p)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	// roll the dice
	t, rs := RollDice(numDice, dieSize, modifier, top)

	fmt.Fprintf(w, "Result: %d<br>%s<br>", t, rs)

}
*/
