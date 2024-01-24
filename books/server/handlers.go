package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
	fmt.Fprintf(w, "Hello there!")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booksArray)
}

func yourSecondMethod(w http.ResponseWriter, r *http.Request) {
	//Your logic
}
func yourThirdMethod(w http.ResponseWriter, r *http.Request) {
	//Your logic
}
func yourFourthMethod(w http.ResponseWriter, r *http.Request) {
	//Your logic
}
