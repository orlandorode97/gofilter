package filter

import (
	"image"
	"image/color"
)

// ApplyNegative applies the negative filter to an image.
func ApplyNegative(img image.Image, e *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			negR, negG, negB := e.GetNegativeRGB(r/EightBits, g/EightBits, b/EightBits)

			output.Set(x, y, color.RGBA{R: negR, G: negG, B: negB, A: Alpha})
		}
	}

	return output
}
