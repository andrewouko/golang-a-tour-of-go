package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tick := time.Tick(1 * time.Second)
	// rem := curr_time.Add(time.Second * 5)
	boom := time.After(5 * time.Second)

	// read from tick and boom channel
	for {
		select {
		case <-tick:
			// rem.Add(time.Second * -1)
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("BOOM!!!")
			return
		default:
			fmt.Println(time.Since(start).Seconds())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
