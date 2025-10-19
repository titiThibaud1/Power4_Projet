package power4

import (
	"html/template"
	"log"
	"net/http"
)

// Route /
func namescreenHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

// Route /difficulty
func difficultyHandler(w http.ResponseWriter, r *http.Request, infosGameVar *InfosGame) {
	if r.FormValue("Player1") != "" && r.FormValue("Player2") != "" {
		infosGameVar.Player1 = r.FormValue("Player1")
		infosGameVar.Player2 = r.FormValue("Player2")
	}

	tmpl, err := template.ParseFiles("pages/difficultyscreen.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// Route /play
func playHandler(w http.ResponseWriter, r *http.Request, infosGameVar *InfosGame) {
	NbTurns = 0
	infosGameVar.CurrentPlayer = 2
	infosGameVar.WaitingPlayer = 1

	if r.FormValue("difficulty") != "" {
		infosGameVar.Difficulty = r.FormValue("difficulty")
	}

	switch infosGameVar.Difficulty {
	case "Easy":
		infosGameVar.Board = [][]int{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		}
		infosGameVar.ArrowLen = 6
		xInit = 6
		yInit = 5
	case "Medium":
		infosGameVar.Board = [][]int{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		}
		infosGameVar.Difficulty = "Medium"
		infosGameVar.ArrowLen = 6
		xInit = 8
		yInit = 5
	case "Hard":
		infosGameVar.Board = [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		}
		infosGameVar.Difficulty = "Hard"
		infosGameVar.ArrowLen = 7
		xInit = 7
		yInit = 6
	}

	tmpl, err := template.ParseFiles("pages/playscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// Route /playTurns
func playTurnsHandler(w http.ResponseWriter, r *http.Request, infosGameVar *InfosGame) {
	NbTurns++
	infosGameVar.ColPlayed = r.FormValue("PlayTurn")
	GamePower4(w, r, NbTurns, infosGameVar)
	tmpl, err := template.ParseFiles("pages/playscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// Route /win
func winHandler(w http.ResponseWriter, infosGameVar *InfosGame) {
	tmpl, err := template.ParseFiles("pages/winscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// Route /draw
func drawHandler(w http.ResponseWriter, infosGameVar *InfosGame) {
	tmpl, err := template.ParseFiles("pages/drawscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}
