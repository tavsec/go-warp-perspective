package go_warp_perspective

import (
	"image"
	"image/color"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func getPixels(src image.Image) ([][]color.RGBA, error) {

	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]color.RGBA
	for y := 0; y < height; y++ {
		var row []color.RGBA
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(src.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) color.RGBA {
	return color.RGBA{R: uint8(r / 257), G: uint8(g / 257), B: uint8(b / 257), A: uint8(a / 257)}
}
