package skinconverter

import (
	"image"
)

func SkinDataToImage(data []byte) (output image.Image) {
	var width, height int
	// all the default skin sizes
	switch len(data) {
	case 64 * 32 * 4:
		width = 64
		height = 32
	case 64 * 64 * 4:
		width = 64
		height = 64
	case 128 * 128 * 4:
		width = 128
		height = 128
	default:
	}
	im := image.NewRGBA(image.Rectangle{
		Max: image.Point{X: width, Y: height},
	})
	im.Pix = data
	output = im
	return
}

func ImageToSkinData(img image.Image) (dat []byte) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			dat = append(dat, uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}
	return dat
}