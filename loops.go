package main

import (
	"fmt"
)

// Uso del único tipo de loop existente: For.
func Sqrt(x float64) float64 {
    z := x/3	// Editable.
	i := 1
	for i < 4 {	// Expresión editable (a mayor nro, mejor aproximación).
		z -= (z*z - x) / (2*z)
		i++
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
