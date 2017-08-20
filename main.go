package main

import (
	"fmt"
	"log"
	"net/http"
)

// shorthand print
func p(s string) { fmt.Println(s) }

func main() {

	p("calling main")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
