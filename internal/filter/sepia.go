package filter

import (
	"image"
	"image/color"
)

// ApplySepia applies the sepia filter to an image.
func ApplySepia(img image.Image, e *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			sepiaR, sepiaG, sepiaB := e.GetSepiaRGB(r/EightBits, g/EightBits, b/EightBits)

			output.Set(x, y, color.RGBA{R: sepiaR, G: sepiaG, B: sepiaB, A: Alpha})
		}
	}

	return output
}
