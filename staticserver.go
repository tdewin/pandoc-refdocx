package main

import (
	"log"
	"net/http"
)

// go run staticserver.go
func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Print("Listening on (:3000) eg http://localhost:3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
