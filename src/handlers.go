package power4

import (
	"html/template"
	"net/http"
)

// Route /
func namescreenHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, gamePower4)
}

// Route /difficulty
func difficultyHandler(w http.ResponseWriter, r *http.Request) {

}

// Route /play
func playHandler(w http.ResponseWriter, r *http.Request) {

}

// Route /win
func winHandler(w http.ResponseWriter, r *http.Request) {

}

// Route /draw
func drawHandler(w http.ResponseWriter, r *http.Request) {

}
