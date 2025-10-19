package power4

import (
	"fmt"
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
		winHandler(w, &infosGameVar)
	})
	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		drawHandler(w, &infosGameVar)
	})

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
