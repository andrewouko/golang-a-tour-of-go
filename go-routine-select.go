package main

import "fmt"

/* func fibonacci(ch chan int, quit chan int) {
	x, y, q := 0, 1, "Quit"
	for {
		select {
		case ch <- x:
			fmt.Println("Run in routine")
			x, y = y, x+y
		case <-quit:
			fmt.Println(q)
			return
		}
	}
}

func main() {
	ch := make(chan int, 10)
	quit := make(chan int)

	go func() {
		for i := 0; i < cap(ch); i++ {
			fmt.Println("Run in main")
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci(ch, quit)
} */

func fibonacci(ch chan int, quit chan string) {
	x, y, q := 0, 1, "Quit"
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case quit <- q:
			// sender closes
			close(ch)
			close(quit)
			return
		}
	}
}

func main() {
	ch := make(chan int, 10)
	quit := make(chan string, 2)

	// sender
	go fibonacci(ch, quit)

	// receiver
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(<-quit)
}
