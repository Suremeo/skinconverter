package skinconverter

import (
	"image"
	"image/color"
)

func SkinDataToImage(data []byte) (output image.Image) {
	var groups [][]byte
	for _, bite := range data {
		if len(groups) == 0 {
			groups = append(groups, []byte{bite})
		} else {
			if len(groups[len(groups) - 1]) == 4 {
				groups = append(groups, []byte{bite})
			} else {
				groups[len(groups) - 1] = append(groups[len(groups) - 1], bite)
			}
		}
	}
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	x, y := 0, 0
	for _, thing := range groups {
		if x == 64 {
			y = y + 1
			x = 0
		}
		img.Set(x, y, color.RGBA{R: thing[0], G: thing[1], B: thing[2], A: thing[3]})
		x = x + 1
	}
	output = img
	return
}

func ImageToSkinData(img image.Image) []byte {
	dat := make([]byte, 64 * 64 * 4)
	x := 0
	y := 0
	for {
		if x == 64 {
			y = y + 1
			x = 0
		}
		if y == 64 {
			break
		}
		here := img.At(x, y)
		r, g, b, a := here.RGBA()
		dat = append(dat, uint8(r), uint8(g), uint8(b), uint8(a))
		x = x + 1
	}
	return dat
}