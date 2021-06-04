package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("port: " + port)
	http.HandleFunc("/health", mainPage)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Health page")
	w.WriteHeader(http.StatusOK)

}
