package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// NOTE: Read populates the given byte slice with data and returns the number of bytes populated and an error value.
// It returns an io.EOF error when the stream ends.

// A reader that produces an infinite stream of char "A"
func (reader MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = byte('A')
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})

	var r MyReader

	b := make([]byte, 3)

	for i := 0; i < 3; i++ {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	s := strings.NewReader("Hello, Reader!")
	for {
		n, err := s.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}
