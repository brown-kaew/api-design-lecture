package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := []byte(`{
			"id": 1,
			"name": "Kaew",
			"age": 18
		}`)
	var u User
	err := json.Unmarshal(data, &u)

	fmt.Printf("type: %T \n", u)
	fmt.Printf("struct: %v \n", u)
	fmt.Println(err)
}
