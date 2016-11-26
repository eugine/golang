package main

import (
	"fmt"
	"io"
	"strings"
)

const rot13A = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const rot13B = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (int, error) {
	n, err := rot13.r.Read(p)
	if err == nil {
		for i := 0; i < n; i++ {
			idx := strings.IndexByte(rot13A, p[i])
			if idx != -1 {
				p[i] = rot13B[idx]
			}
		}
	}
	return 0, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	// io.Copy(os.Stdout, &r)
	b := make([]byte, 22)
	r.Read(b)
	fmt.Printf(string(b) + "\n")
}
