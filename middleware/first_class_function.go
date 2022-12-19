package main

import "fmt"

type Math func(int, int) int

func cal(sn Math) int {
	return sn(5, 4)
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fn := sum

	r1 := fn(1, 2)
	fmt.Println("fn(1, 2):", r1) //fn(1, 2): 3

	r2 := cal(fn)
	fmt.Println("cal(fn):", r2) //cal(fn): 9

	r3 := cal(sum)
	fmt.Println("cal(sum):", r3) //cal(sum): 9
}
