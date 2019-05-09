package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(ch byte) byte {
	if (ch >= 'A' && ch <= 'M') || (ch >= 'a' && ch <= 'm') {
		return ch + 13
	}
	if (ch >= 'N' && ch <= 'Z') || (ch >= 'n' && ch <= 'z') {
		return ch - 13
	}
	return ch
}

func (rot13_r rot13Reader) Read(b []byte) (n int, err error) {
	t := make([]byte, len(b))
	n, err = rot13_r.r.Read(t)
	for i := 0; i < n; i++ {
		b[i] = rot13(t[i]) // 解码
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
