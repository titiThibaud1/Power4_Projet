package power4

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type InfosGame struct {
	Player1       string
	Player2       string
	CurrentPlayer string
	WaitingPlayer string
	Difficulty    string
	ArrowLen      int
	Board         [][]int
}

var NbTurns int
var infosGameVar InfosGame

func Server() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", namescreenHandler)
	http.HandleFunc("/difficulty", func(w http.ResponseWriter, r *http.Request) {
		difficultyHandler(w, r, &infosGameVar)
	})
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		playHandler(w, r, &infosGameVar)
	})
	http.HandleFunc("/playTurns", func(w http.ResponseWriter, r *http.Request) {
		playTurnsHandler(w, r, &infosGameVar)
	})
	http.HandleFunc("/win", winHandler)
	http.HandleFunc("/draw", drawHandler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ======================================================================================

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
	infosGameVar.Player1 = r.FormValue("Player1")
	infosGameVar.Player2 = r.FormValue("Player2")

	tmpl, err := template.ParseFiles("pages/difficultyscreen.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// Route /play
func playHandler(w http.ResponseWriter, r *http.Request, infosGameVar *InfosGame) {
	NbTurns = 0
	infosGameVar.Difficulty = r.FormValue("difficulty")
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
	}

	tmpl, err := template.ParseFiles("pages/playscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

func playTurnsHandler(w http.ResponseWriter, r *http.Request, infosGameVar *InfosGame) {
	NbTurns++
	GamePower4(NbTurns, infosGameVar)
	tmpl, err := template.ParseFiles("pages/playscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// Route /win
func winHandler(w http.ResponseWriter, r *http.Request) {

}

// Route /draw
func drawHandler(w http.ResponseWriter, r *http.Request) {

}

// ==================================================================================

func NewGame() {

}

func GamePower4(NbTurn int, infosGameVar *InfosGame) {
	if NbTurn%2 != 0 {
		infosGameVar.CurrentPlayer = "Player1"
		infosGameVar.WaitingPlayer = "Player2"
	} else {
		infosGameVar.CurrentPlayer = "Player2"
		infosGameVar.WaitingPlayer = "Player1"
	}

}

//board[1][2] = 42
