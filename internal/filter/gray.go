package filter

import (
	"image"
	"image/color"
)

// ApplyGray applies the gray scale filter to an image.
func ApplyGray(img image.Image, e *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := e.GetGrayRGB(r/EightBits, g/EightBits, b/EightBits)

			output.Set(x, y, color.RGBA{R: gray, G: gray, B: gray, A: Alpha})
		}
	}

	return output
}
