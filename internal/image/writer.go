package image

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/OrlandoRomo/go-filter/internal/filter"
)

// SaveImage saves an image with the same format as the input.
func SaveImage(img image.Image, outputPath, inputExt string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	switch inputExt {
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, nil)
	default:
		return png.Encode(file, img)
	}
}

// ProcessAndSave reads an image, applies a filter, and saves the result.
func ProcessAndSave(inputPath, outputDir string, filterFunc filter.FilterFunc, progress chan<- float64) (string, error) {
	file, err := ReadFile(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Reset file for potential re-read.
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	// Apply filter.
	effect := &filter.Effect{}
	filtered := filterFunc(img, effect)

	// Report progress at 50% after filter is applied.
	if progress != nil {
		progress <- 0.5
	}

	// Generate output path.
	randomName := filter.RandomName()
	outputPath := GenerateOutputPath(outputDir, inputPath, randomName)

	// Save image.
	if err := SaveImage(filtered, outputPath, filter.GetExtension(inputPath)); err != nil {
		return "", err
	}

	// Report progress complete.
	if progress != nil {
		progress <- 1.0
	}

	return outputPath, nil
}
