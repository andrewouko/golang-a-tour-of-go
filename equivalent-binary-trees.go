package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// close channel before function returns
	defer close(ch)
	var traverse func(t *tree.Tree)
	traverse = func(t *tree.Tree) {
		if t.Left != nil {
			traverse(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			traverse(t.Right)
		}
		/* // return function once tree being
		// traversed has reached nil
		if t == nil {
			return
		}
		// walk left tree recursively
		Walk(t.Left, ch)
		// add value from left
		ch <- t.Value
		Walk(t.Right, ch) */
	}
	traverse(t)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t1, ch2)
	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true

}

func main() {
	tree1 := tree.New(1)
	tree2 := tree.New(1)

	fmt.Println(Same(tree1, tree2))
}
