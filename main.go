package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

// shorthand print
func p(s string) { fmt.Println(s) }

func init() {
	log.Println("init called...")

}

func main() {

	p("calling main")

	router := NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))

	//log.Fatal(http.ListenAndServe(":8080", router))
}
