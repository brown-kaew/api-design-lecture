package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Kaew", Age: 18},
	{ID: 2, Name: "Ewka", Age: 22},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(b)
		return
	}

	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var u User
		err = json.Unmarshal(b, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		users = append(users, u)
		// w.Write([]byte(`hello POST create user`))
		fmt.Fprintf(w, "hello %s create user", r.Method)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`OK`))
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Server started ad : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("bye bye")
}
