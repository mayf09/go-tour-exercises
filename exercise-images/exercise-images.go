package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image 无意义的 Image 实现
type Image struct {
	w, h int
	v    uint8
}

// TODO: 理解实际的 Image

// ColorModel .
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds .
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h) // FIXME
}

// At .
func (img Image) At(x, y int) color.Color {
	return color.RGBA{img.v, img.v, 255, 255} // FIXME
}

func main() {
	m := Image{100, 100, 0}
	pic.ShowImage(m)
}
