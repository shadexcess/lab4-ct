package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
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

	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", helloHandler)

	log.Printf("Server is starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
