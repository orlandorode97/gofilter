package filter

import (
	"image"
	"math/rand"
	"path/filepath"
)

// Effect implements filter operations.
type Effect struct{}

func init() {
	// Go 1.20+ auto-seeds the random generator.
	// rand.Seed is deprecated but kept for backward compatibility.
	_ = rand.Intn
}

// GetGrayRGB returns RGB values that represent the gray scale tone.
func (e *Effect) GetGrayRGB(r, g, b uint32) uint8 {
	return uint8((RedWaveLength * float64(r/EightBits)) + (GreenWaveLength * float64(g/EightBits)) + (BlueWaveLength * float64(b/EightBits)))
}

// GetSepiaRGB gets the RGB values that represent the sepia scale tone.
func (e *Effect) GetSepiaRGB(r, g, b uint32) (uint8, uint8, uint8) {
	tr := (SepiaRedForRed * float64(r)) + (SepiaGreenForRed * float64(g)) + (SepiaBlueForRed * float64(b))
	tg := (SepiaRedForGreen * float64(r)) + (SepiaGreenForGreen * float64(g)) + (SepiaBlueForGreen * float64(b))
	tb := (SepiaRedForBlue * float64(r)) + (SepiaGreenForBlue * float64(g)) + (SepiaBlueForBlue * float64(b))

	if int(tr) > RGBA {
		tr = float64(RGBA)
	}

	if int(tg) > RGBA {
		tg = float64(RGBA)
	}

	if int(tb) > RGBA {
		tb = float64(RGBA)
	}

	return uint8(tr), uint8(tg), uint8(tb)
}

// GetNegativeRGB returns the negative RGB values.
func (e *Effect) GetNegativeRGB(r, g, b uint32) (uint8, uint8, uint8) {
	return uint8(MaxIntensity - r), uint8(MaxIntensity - g), uint8(MaxIntensity - b)
}

// GetSketchRGB returns the sketch RGB values based on intensity.
func (e *Effect) GetSketchRGB(r, g, b uint32) (uint8, uint8, uint8) {
	intensity := e.GetGrayRGB(r, g, b)
	if intensity > IntensityFactor {
		return HighestValue, HighestValue, HighestValue
	}

	if intensity > 100 {
		return MeanValue, MeanValue, MeanValue
	}

	return LowestValue, LowestValue, LowestValue
}

// GetRedRGB returns the red scale RGB values.
func (e *Effect) GetRedRGB(_, g, b uint32) (uint8, uint8, uint8) {
	green := g
	blue := b

	if green > MinGreenForRed {
		green = MinGreenForRed
	}

	if blue > MinBlueForRed {
		blue = MinBlueForRed
	}

	return uint8(MaxRed), uint8(green), uint8(blue)
}

// GetGreenRGB returns the green scale RGB values.
func (e *Effect) GetGreenRGB(r, g, b uint32) (uint8, uint8, uint8) {
	red := r
	green := g
	blue := b

	if red > MinRedForGreen {
		red = MinRedForGreen
	}

	if green < MaxGreen {
		green = MaxGreen
	}

	if blue > MinBlueForGreen {
		blue = MinBlueForGreen
	}

	return uint8(red), uint8(green), uint8(blue)
}

// GetBlueRGB returns the blue scale RGB values.
func (e *Effect) GetBlueRGB(r, g, b uint32) (uint8, uint8, uint8) {
	red := r
	green := g
	blue := b

	if blue < MaxBlue {
		blue = MaxBlue
	}

	if red > MinRedForBlue {
		red = MinRedForBlue
	}

	if green > MinGreenForBlue {
		green = MinGreenForBlue
	}

	return uint8(red), uint8(green), uint8(blue)
}

// RandomName generates a random name for output files.
func RandomName() string {
	b := make([]byte, MaxCharacters)
	for i := range b {
		b[i] = LetterRunes[rand.Intn(len(LetterRunes))]
	}
	return string(b)
}

// SupportedExtensions returns the map of supported image extensions.
func SupportedExtensions() map[string]bool {
	return map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
}

// IsValidExtension checks if the file extension is supported.
func IsValidExtension(name string) bool {
	extension := filepath.Ext(name)
	supported := SupportedExtensions()
	_, ok := supported[extension]

	return ok
}

// GetExtension returns the file extension from a path.
func GetExtension(filePath string) string {
	return filepath.Ext(filePath)
}

// FilterFunc defines a function type for applying filters to an image.
type FilterFunc func(img image.Image, effect *Effect) image.Image

// FilterInfo holds metadata about a filter.
type FilterInfo struct {
	Name        string
	Description string
	Apply       FilterFunc
}
