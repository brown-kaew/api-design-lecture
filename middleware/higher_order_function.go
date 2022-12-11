package main

import "fmt"

type Decorator func(s string) error

func Hello(s string) error {
	fmt.Println("hello", s)
	return nil
}

func Use(next Decorator) Decorator {
	return func(s string) error {
		fmt.Println("do somthing before")
		r := s + " from the other side"
		return next(r)
	}
}

func main() {
	wrapped := Use(Hello)
	w := wrapped("world")
	fmt.Println("result:", w)
}
