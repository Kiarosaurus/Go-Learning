// https://go-tour-lat.appspot.com/moretypes/18

package main

import "golang.org/x/tour/pic"

// Slice de Slices.
// (make para hacer slices. No podemos hacer arrays en tiempo de compilación).
func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	
	for i:=0; i<dy; i++ {
		img[i] = make([]uint8, dx)
	}
	return img
}


func main() {
	pic.Show(Pic)
}
