package filter

import (
	"image"
	"image/color"
)

// ApplyRed applies the red scale filter to an image.
func ApplyRed(img image.Image, e *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			redR, redG, redB := e.GetRedRGB(r/EightBits, g/EightBits, b/EightBits)

			output.Set(x, y, color.RGBA{R: redR, G: redG, B: redB, A: Alpha})
		}
	}

	return output
}
