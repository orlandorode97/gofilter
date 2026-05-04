package filter

import (
	"image"
	"image/color"
)

// ApplySketch applies the sketch filter to an image.
func ApplySketch(img image.Image, e *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			sketchR, sketchG, sketchB := e.GetSketchRGB(r/EightBits, g/EightBits, b/EightBits)

			output.Set(x, y, color.RGBA{R: sketchR, G: sketchG, B: sketchB, A: Alpha})
		}
	}

	return output
}
