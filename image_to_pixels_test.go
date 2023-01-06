package go_warp_perspective

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestGetPixels(t *testing.T) {
	width, height := 200, 200
	testColorRed := color.RGBA{R: 200, A: 255}
	testColorBlue := color.RGBA{B: 200, A: 255}
	testColorWhite := color.RGBA{}
	src := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: width, Y: height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				src.Set(x, y, testColorRed)
			case x >= width/2 && y >= height/2: // lower right quadrant
				src.Set(x, y, testColorBlue)
			default:
			}
		}
	}

	pixels, err := GetPixels(src)
	if err != nil {
		t.Fatal("GetPixels returned error: " + err.Error())
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				if pixels[x][y] != testColorRed {
					t.Fatal(fmt.Sprintf("Color at [%d, %d] should be (%d, %d, %d, %d), but is (%d, %d, %d, %d)", x, y, testColorRed.R, testColorRed.G, testColorRed.B, testColorRed.A, pixels[x][y].R, pixels[x][y].G, pixels[x][y].B, pixels[x][y].A))
				}
			case x >= width/2 && y >= height/2: // lower right quadrant
				if pixels[x][y] != testColorBlue {
					t.Fatal(fmt.Sprintf("Color at [%d, %d] should be (%d, %d, %d, %d), but is (%d, %d, %d, %d)", x, y, testColorBlue.R, testColorBlue.G, testColorBlue.B, testColorBlue.A, pixels[x][y].R, pixels[x][y].G, pixels[x][y].B, pixels[x][y].A))
				}
			default:
				if pixels[x][y] != testColorWhite {
					t.Fatal(fmt.Sprintf("Color at [%d, %d] should be (%d, %d, %d, %d), but is (%d, %d, %d, %d)", x, y, testColorWhite.R, testColorWhite.G, testColorWhite.B, testColorWhite.A, pixels[x][y].R, pixels[x][y].G, pixels[x][y].B, pixels[x][y].A))
				}
			}
		}
	}
}
