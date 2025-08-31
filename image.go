package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

// Definiendo mi Image que cumple con la interface (image.Image)
type Image struct{
	w,h int
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func (im Image) At(x, y int) color.Color {
	v := uint8((x*y)) // patrón sencillo
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 100}  // probando
	pic.ShowImage(m)
}
