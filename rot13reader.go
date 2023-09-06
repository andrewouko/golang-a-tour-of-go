package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)

	if err != nil {
		return 0, err
	}

	for i := 0; i < len(b); i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		}
		if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'N' && b[i] <= 'Z') {
			b[i] -= 13
		}
	}
	return len(b), nil
}

func main() {
	reader := strings.NewReader("The lazy brown fox jumped over the fence")
	rot13 := &rot13Reader{reader}

	// read from rot13Reader
	b := make([]byte, 10)
	for {
		n, err := rot13.Read(b)
		fmt.Print(string(b[:n]))
		if err == io.EOF {
			break
		}
	}

	// write to stdout from what is read from ro13.Read()
	io.Copy(os.Stdout, rot13)
}
