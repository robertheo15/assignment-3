package main

import (
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", Randomize)

	log.Println("Server started at port", PORT)
	http.ListenAndServe(PORT, nil)
}
