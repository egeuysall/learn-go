package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Define a simple struct
	type Note struct {
		Text string `json:"text"`
	}

	// Struct to JSON
	note := Note{Text: "Hello JSON"}
	jsonData, _ := json.Marshal(note)
	fmt.Println(string(jsonData)) // Output: {"text":"Hello JSON"}

	// JSON to struct
	input := `{"text":"Decoded!"}`
	var decodedNote Note
	_ = json.Unmarshal([]byte(input), &decodedNote)
	fmt.Println(decodedNote.Text) // Output: Decoded!
}
