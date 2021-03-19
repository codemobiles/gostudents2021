package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./build")) // New code
	http.Handle("/", fileServer)                        // New code

	fmt.Printf("Starting server at port 5001\n")
	if err := http.ListenAndServe(":5001", nil); err != nil {
		log.Fatal(err)
	}
}
