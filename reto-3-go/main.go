package main

import (
	"fmt"
	"net/http"
	"time"
)

var userVisits = make(map[string]int)

func trackVisitHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "User parameter is missing", http.StatusBadRequest)
		return
	}

	time.Sleep(500 * time.Millisecond)

	userVisits[username]++
	visits := userVisits[username]

	fmt.Fprintf(w, "User %s has visited %d times\n", username, visits)
}

func main() {
	http.HandleFunc("/visit", trackVisitHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
