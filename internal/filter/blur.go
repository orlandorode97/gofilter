package filter

import (
	"image"
	"image/color"
)

// ApplyBlur applies the blur filter to an image.
func ApplyBlur(img image.Image, _ *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
		for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
			var r, g, b uint32

			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					pr, pg, pb, _ := img.At(x+dx, y+dy).RGBA()
					r += pr
					g += pg
					b += pb
				}
			}

			output.Set(x, y, color.RGBA{
				R: uint8(r / 9 >> 8),
				G: uint8(g / 9 >> 8),
				B: uint8(b / 9 >> 8),
				A: Alpha,
			})
		}
	}

	return output
}
