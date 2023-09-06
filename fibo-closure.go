package main

import "fmt"

func fibonnaci() func() int {
	x, y := 0, 1
	return func() int {
		res := x
		x, y = y, x+y
		return res
	}
}

func main() {
	num := fibonnaci()
	for i := 0; i < 10; i++ {
		fmt.Println(num())
	}
}
