package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the MainPage")
	w.WriteHeader(http.StatusOK)

}
