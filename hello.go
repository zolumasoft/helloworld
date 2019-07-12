package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

// F sends greeting
func F(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Go Function"))
}

func main() {
	http.HandleFunc("/", F)
	log.Printf("Server started @ %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
