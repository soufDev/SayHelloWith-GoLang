package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloTeam(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hey GoLand I Love YOU ALL")
}
func main() {
	http.HandleFunc("/", sayHelloTeam)
	log.Fatal(http.ListenAndServe(":8000", nil))
}