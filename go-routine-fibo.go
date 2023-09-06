package main

import "fmt"

func fibonacci(ch chan int) {
	x, y := 0, 1
	for i := 0; i < cap(ch); i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
