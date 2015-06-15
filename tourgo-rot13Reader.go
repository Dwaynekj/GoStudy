package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotter *rot13Reader) Read(b []byte) (int, error) {
	buf := make([]byte, len(b))
	n, err := rotter.r.Read(buf)
	for i, v := range buf {
		b[i] = v + 13
		//fmt.Sprintf("b[:n] = %q\n", b[:n])
	}
	return n, err
	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
