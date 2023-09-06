package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func half(num int) (int, bool) {
	half := num / 2
	if half%2 == 0 {
		return half, true
	} else {
		return half, false
	}
}

func max(args ...int) int {
	// smallest possible number
	/* max := int(math.Inf(-1))
	for _, v := range args {
		if v > max {
			max = v
		}
	} */
	sort.Ints(args)
	return args[len(args)-1]

}

func makeOddGenerator() func() int {
	x := 1
	return func() int {
		res := x
		x += 2
		return res
	}
}

func fib(n uint) uint {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func testError() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("Testing a panic recovery")
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	arr := make([]int, 10)
	for i := 1; i < len(arr); i++ {
		arr[i] = r.Intn(len(arr))
	}
	fmt.Printf("sum of %v is %v \n", arr, sum(arr))
	num := arr[r.Intn(len(arr))]
	half, even := half(num)
	fmt.Printf("Num: %v Half: %v Even: %v  \n", num, half, even)
	max := max(arr...)
	fmt.Printf("Max num is %v \n", max)
	fmt.Printf("Generating %v odd nums \n", max)
	nextOdd := makeOddGenerator()
	for i := 0; i < max; i++ {
		fmt.Println(nextOdd())
	}
	fmt.Printf("Fibonacci num %v is %v \n", num, fib(uint(num)))
	testError()
}
