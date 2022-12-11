package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"name": "Kaew"}`))
	})

	log.Println("Server started ad : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("bye bye")
}
