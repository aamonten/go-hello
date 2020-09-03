package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aamonten/go-hello/morestrings"
)

func reverseHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, morestrings.ReverseRunes("Hello, world!"))

}

func main() {

	http.HandleFunc("/", reverseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
