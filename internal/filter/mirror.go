package filter

import (
	"image"
	"image/color"
)

// ApplyMirror applies the mirror effect to an image.
func ApplyMirror(img image.Image, _ *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	width := bounds.Max.X

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < width; x++ {
			r, g, b, a := img.At(width-1-x, y).RGBA()

			output.Set(x, y, color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			})
		}
	}

	return output
}
