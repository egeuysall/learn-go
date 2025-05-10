package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Highlight struct {
	Text     string   `json:"text"`
	Mood     int      `json:"mood"`
	Category []string `json:"category"`
	Date     string   `json:"date"`
	ID       int      `json:"id"`
}

var highlights = []Highlight{
	{
		ID:       1,
		Text:     "Built my first Go API!",
		Mood:     5,
		Category: []string{"coding", "learning"},
		Date:     "2025-05-10",
	},
	{
		ID:       2,
		Text:     "Explored Go slices and handlers.",
		Mood:     4,
		Category: []string{"go", "backend"},
		Date:     "2025-05-11",
	},
	{
		ID:       3,
		Text:     "Felt stuck debugging a routing issue.",
		Mood:     2,
		Category: []string{"debugging", "frustration"},
		Date:     "2025-05-12",
	},
	{
		ID:       4,
		Text:     "Had a breakthrough with Docker containers.",
		Mood:     5,
		Category: []string{"docker", "deployment"},
		Date:     "2025-05-13",
	},
	{
		ID:       5,
		Text:     "Struggled with Go error handling patterns.",
		Mood:     3,
		Category: []string{"go", "learning"},
		Date:     "2025-05-14",
	},
	{
		ID:       6,
		Text:     "Successfully connected frontend to Go API!",
		Mood:     5,
		Category: []string{"fullstack", "milestone"},
		Date:     "2025-05-15",
	},
}

func handleHighlight(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var newHighlight Highlight

		err := json.NewDecoder(r.Body).Decode(&newHighlight)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		newHighlight.ID = len(highlights) + 1
		highlights = append(highlights, newHighlight)

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(newHighlight)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(highlights)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Only POST and GET methods are allowed", http.StatusMethodNotAllowed)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// CORS for root too
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "Welcome to the API! Please go to /highlight to post a new highlight.")
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/highlight", handleHighlight)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
