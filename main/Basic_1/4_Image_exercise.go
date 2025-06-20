package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	// v := uint8((x + y) / 2)   // Một gradient chéo
	// v := uint8(x * y)        // Các đường cong hyperbol
	// v := uint8(x ^ y)        // Một mẫu kẻ sọc

	v := uint8((x ^ y))
	return color.RGBA{v, 255, v, 255}
}

func main() {
	m := Image{width: 256, height: 256}
	pic.ShowImage(m)
}
