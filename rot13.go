package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)	// Leemos la lectura (está comprimido doblemente).
	for i := 0; i<n; i++ {
		b := p[i]
		
		//rot13
		if b >= 'A' && b <= 'Z' {		// MAYÚSCULAS
			p[i] = 'A' + (b-'A'+13)%26
		} else if b >= 'a' && b <= 'z' { // minúsculas
			p[i] = 'a' + (b-'a'+13)%26
		}	// default: dejamos igual
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
