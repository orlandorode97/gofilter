package filter

import (
	"image"
	"image/color"
)

// ApplyBlue applies the blue scale filter to an image.
func ApplyBlue(img image.Image, e *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			blueR, blueG, blueB := e.GetBlueRGB(r/EightBits, g/EightBits, b/EightBits)

			output.Set(x, y, color.RGBA{R: blueR, G: blueG, B: blueB, A: Alpha})
		}
	}

	return output
}
