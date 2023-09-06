package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx int, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := range res {
		res[y] = make([]uint8, dx)
		for x := range res[y] {
			val := (x + y) / 2
			// val := x*y
			// val := x^y
			res[y][x] = uint8(val)
		}
	}
	return res
}

func main() {
	// fmt.Println(Pic(6, 6))
	pic.Show(Pic)
}
