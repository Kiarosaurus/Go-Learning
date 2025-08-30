package main

import "fmt"

// fibonacci es una función que devuelve un int.
// Usando switch y slices.
func fibonacci() func() int {
	a := make([]int, 0)

	return func() int {
		switch len(a) {
		case 0:
			a = append(a, 0)
			return 0
		case 1:
			a = append(a, 1)
			return 1
		default:
			aux := a[1]
			a[1] = a[0] + a[1]
			a[0] = aux
			return a[1]
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
