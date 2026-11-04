package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RequestPayload struct {
	Name string `json:"text"`
}

func countWordsHandler(w http.ResponseWriter, r *http.Request) {
	var d RequestPayload
	_ = json.NewDecoder(r.Body).Decode(&d)

	if d.Text == "" {
		d.Text = r.URL.Query().Get("text")
	}

	if d.Text == "" {
		http.Error(w, "Attach text!", http.StatusBadRequest)
		return
	}

	words := strings.Fields(d.Text)
	count := len(words)

	fmt.Fprint(w, "Words number: ", count)
}

func main() {
	port := "8080"
	http.HandleFunc("/", countWordsHandler)
	_ = http.ListenAndServe(":"+port, nil)
}
