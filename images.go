package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
	v    uint8
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *Image) At(x, y int) color.Color {
	// val := (x + y) / 2
	val := x * y
	// val := x^y
	return color.RGBA{img.v + uint8(val), img.v + uint8(val), 255, 255}
}

func main() {
	img := &Image{100, 100, 128}
	pic.ShowImage(img)
}
