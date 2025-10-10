package main

import (
	"net/http"
	power4 "power4/src"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	power4.Server()
}
