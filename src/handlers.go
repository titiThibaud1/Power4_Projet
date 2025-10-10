package power4

import "net/http"

func Handlers() {
	http.HandleFunc("/", namescreenHandler)
	http.HandleFunc("/difficulty", difficultyHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/win", winHandler)
	http.HandleFunc("/draw", drawHandler)
}
