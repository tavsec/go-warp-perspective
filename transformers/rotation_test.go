package transformers

import (
	go_warp_perspective "github.com/tavsec/go-warp-perspective"
	"image"
	"image/color"
	"math"
	"testing"
)

func TestSimpleRotation(t *testing.T) {
	img := createTestImage()
	imgPixels, _ := go_warp_perspective.GetPixels(img)
	simpleRotationByAngle(math.Pi, &imgPixels)
}

func createTestImage() image.Image {
	width := 200
	height := 100

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
			}
		}
	}
	return img
}
