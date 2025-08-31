package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Agregar un metodo Read([]byte) (int, error) a MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
	for i := range b {  // Para asegurar el "flujo infinito" de As.
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
