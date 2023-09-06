package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (reader MyReader) Read(b []byte) (n int, err error) {
	err = nil
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	n = len(b)
	return
}

func main() {
	r := MyReader{}
	b := make([]byte, 8)
	n, err := r.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < n; i++ {
		fmt.Println(string(b[i]))
	}
	reader.Validate(MyReader{})
}
