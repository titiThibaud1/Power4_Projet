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
	CurrentPlayer int
	WaitingPlayer int
	Winner        string
	ColPlayed     string
	Difficulty    string
	ArrowLen      int
	Board         [][]int
}

var NbTurns, y, xInit, yInit int
var infosGameVar InfosGame

func Server() {
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
	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		winHandler(w, r, &infosGameVar)
	})
	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		winHandler(w, r, &infosGameVar)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
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
	infosGameVar.CurrentPlayer = 2
	infosGameVar.WaitingPlayer = 1
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
func winHandler(w http.ResponseWriter, r *http.Request, infosGameVar *InfosGame) {
	tmpl, err := template.ParseFiles("pages/winscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// Route /draw
func drawHandler(w http.ResponseWriter, r *http.Request, infosGameVar *InfosGame) {
	tmpl, err := template.ParseFiles("pages/drawscreen.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, infosGameVar)
}

// ==================================================================================

func NewGame() {

}

func GamePower4(w http.ResponseWriter, r *http.Request, NbTurn int, infosGameVar *InfosGame) {
	if NbTurn%2 != 0 {
		infosGameVar.CurrentPlayer = 1
		infosGameVar.WaitingPlayer = 2
	} else {
		infosGameVar.CurrentPlayer = 2
		infosGameVar.WaitingPlayer = 1
	}

	y = 0

	switch infosGameVar.ColPlayed {
	case "Col0":
		y = 0
	case "Col1":
		y = 1
	case "Col2":
		y = 2
	case "Col3":
		y = 3
	case "Col4":
		y = 4
	case "Col5":
		y = 5
	case "Col6":
		y = 6
	}

	var x int
	for x = xInit; x >= 0; x-- {
		if infosGameVar.Board[x][y] == 0 {
			infosGameVar.Board[x][y] = infosGameVar.CurrentPlayer
			break
		}
	}
	checkWin(w, r, infosGameVar)

}

func checkWin(w http.ResponseWriter, r *http.Request, infoGameVar *InfosGame) {
	var VerifVal int
	var VerifDraw bool

	for i := 0; i <= xInit; i++ {
		for j := 0; j <= yInit; j++ {
			VerifVal = infoGameVar.Board[i][j]

			if j+3 <= yInit && //Verif horizontal
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i][j+1] &&
				VerifVal == infoGameVar.Board[i][j+2] &&
				VerifVal == infoGameVar.Board[i][j+3] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/win", http.StatusSeeOther)
			}

			if i+3 <= xInit && //Verif vertical
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i+1][j] &&
				VerifVal == infoGameVar.Board[i+2][j] &&
				VerifVal == infoGameVar.Board[i+3][j] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/win", http.StatusSeeOther)
			}

			if i-3 >= 0 && j-3 >= 0 && //Verif diago bas droit --> haut gauche
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i-1][j-1] &&
				VerifVal == infoGameVar.Board[i-2][j-2] &&
				VerifVal == infoGameVar.Board[i-3][j-3] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/win", http.StatusSeeOther)
			}

			if i-3 >= 0 && j+3 <= yInit && //Verif diago bas gauche --> haut droite
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i-1][j+1] &&
				VerifVal == infoGameVar.Board[i-2][j+2] &&
				VerifVal == infoGameVar.Board[i-3][j+3] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/win", http.StatusSeeOther)
			}
		}
	}
	for CptDraw := 0; CptDraw <= yInit; CptDraw++ {
		if infoGameVar.Board[0][CptDraw] != 0 {
			VerifDraw = true
		} else {
			VerifDraw = false
			break
		}
	}
	if VerifDraw {
		http.Redirect(w, r, "/draw", http.StatusSeeOther)
	}
}

//board[1][2] = 42
