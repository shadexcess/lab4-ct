package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type RequestPayload struct {
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var d RequestPayload
	_ = json.NewDecoder(r.Body).Decode(&d)

	if d.Name == "" {
		d.Name = r.URL.Query().Get("name")
	}

	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}

	fmt.Fprint(w, "Hello, ", html.EscapeString(d.Name), "!")
}

func main() {
	port := "8080"
	http.HandleFunc("/", helloHandler)
	_ = http.ListenAndServe(":"+port, nil)
}
