package power4

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", namescreenHandler)
	http.HandleFunc("/difficulty", difficultyHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/win", winHandler)
	http.HandleFunc("/draw", drawHandler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
