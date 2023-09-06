package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func createNewList[T string](value T) *List[T] {
	return &List[T]{next: nil, val: value}
}

/* func createNewList(value string) *List[string] {
	return &List[string]{next: nil, val: value}
} */

func main() {
}
