package main

import (
	"fmt"
)

// Añadiendo el tipo error (estructura).
type ErrNegativeSqrt struct {
	numero float64
	
}

// Añadiendo el método del tipo error.
func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("No se puede calcular la raíz cuadrada de un número negativo: %v", e.numero)
}

// Función Sqrt editada por si hay error.
func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		z := x/3	// Editable.
		i := 1
		for i < 4 {	// Expresión editable (a mayor nro, mejor aproximación).
			z -= (z*z - x) / (2*z)
			i++
		}
		return z, nil
	} else {
		return -1.0, &ErrNegativeSqrt{x}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
