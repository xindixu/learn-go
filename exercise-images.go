package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	model color.Model
}

func (img Image) ColorModel() color.Model {
	return img.model
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 300, 300)
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) * 2)
	return color.RGBA{v, v, 255, 255}
}
func main() {
	m := Image{color.RGBAModel}
	pic.ShowImage(m)
}
