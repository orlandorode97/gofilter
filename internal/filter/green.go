package filter

import (
	"image"
	"image/color"
)

// ApplyGreen applies the green scale filter to an image.
func ApplyGreen(img image.Image, e *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			greenR, greenG, greenB := e.GetGreenRGB(r/EightBits, g/EightBits, b/EightBits)

			output.Set(x, y, color.RGBA{R: greenR, G: greenG, B: greenB, A: Alpha})
		}
	}

	return output
}
