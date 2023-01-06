package transformers

import (
	"image/color"
	"math"
	"sync"
)

func simpleRotationByAngle(angle float64, srcPixels *[][]color.RGBA) {
	ppixels := *srcPixels
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	newImage, shiftX, shiftY := emptyImageAndShift(ppixels, cos, sin)
	var wg sync.WaitGroup
	for i := 0; i < len(ppixels); i++ {
		for j := 0; j < len(ppixels[i]); j++ {
			wg.Add(1)
			go func(i, j int) {
				xNew := int(float64(i)*cos-float64(j)*sin) + shiftX
				yNew := int(float64(i)*sin+float64(j)*cos) + shiftY
				newImage[xNew][yNew] = ppixels[i][j]
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	*srcPixels = newImage
}

func emptyImageAndShift(ppixels [][]color.RGBA, cos, sin float64) (img [][]color.RGBA, shiftX, shiftY int) {
	var lock = sync.Mutex{}
	var wg sync.WaitGroup
	var xMax, yMax, xLow, yLow int
	for i := 0; i < len(ppixels); i++ {
		for j := 0; j < len(ppixels[i]); j++ {
			wg.Add(1)
			go func(i, j int) {
				xNew := int(float64(i)*cos - float64(j)*sin)
				yNew := int(float64(i)*sin + float64(j)*cos)
				if xNew > xMax {
					lock.Lock()
					xMax = xNew
					lock.Unlock()
				}
				if yNew > yMax {
					lock.Lock()
					yMax = yNew
					lock.Unlock()
				}
				if xNew < xLow {
					lock.Lock()
					xLow = xNew
					lock.Unlock()
				}
				if yNew < yLow {
					lock.Lock()
					yLow = yNew
					lock.Unlock()
				}
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	shiftX = int(math.Abs(float64(xLow)))
	width := xMax + shiftX
	shiftY = int(math.Abs(float64(yLow)))
	height := yMax + shiftY
	newImage := make([][]color.RGBA, width+1)
	for i := 0; i < len(newImage); i++ {
		newImage[i] = make([]color.RGBA, height+1)
	}
	return newImage, shiftX, shiftY
}
