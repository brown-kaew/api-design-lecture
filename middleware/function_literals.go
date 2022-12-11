package main

import "fmt"

func main() {
	r := func(a, b int) bool {
		return a < b
	}(3, 4)
	fmt.Println("result:", r) //result: true
}
