package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home functions that handle web request need to handle two parameters a responseWriter and request
func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//start a webserver in go
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
