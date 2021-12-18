package main

import (
	"fmt"
	"log"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func main() {
	http.HandleFunc("/", homePageHandler)
	var port = 3000
	fmt.Println("Server running on port", port)
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
