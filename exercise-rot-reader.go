package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const upperACode = 65
const lowerACode = 97

func rot13(b byte) byte {
	const rotation = 13
	const total = 26
	if b >= 'A' && b < 'Z' {
		return upperACode + (b-upperACode+rotation)%total
	} else if b >= 'a' && b < 'z' {
		return lowerACode + (b-lowerACode+rotation)%total
	}
	return b

}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	num, err := rr.r.Read(b)

	if err != nil {
		return num, err
	}

	for i := 0; i < num; i++ {
		b[i] = rot13(b[i])
	}
	return num, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
