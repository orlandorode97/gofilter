package image

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/OrlandoRomo/go-filter/internal/filter"
)

// ErrInvalidExtension is returned when the file extension is not supported.
var ErrInvalidExtension = errors.New("file is not supported: supported extensions are .png, .jpg, .jpeg")

// ErrNoFileProvided is returned when no file path is provided.
var ErrNoFileProvided = errors.New("command requires a path as argument")

// ReadFile opens and validates an image file.
func ReadFile(filePath string) (*os.File, error) {
	if len(filePath) == 0 {
		return nil, ErrNoFileProvided
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	if !filter.IsValidExtension(file.Name()) {
		file.Close()

		return nil, ErrInvalidExtension
	}

	return file, nil
}

// GetExtension returns the file extension.
func GetExtension(filePath string) string {
	return filepath.Ext(filePath)
}

// ValidateOutputPath checks if the output directory exists and is writable.
func ValidateOutputPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("output path does not exist")
		}

		return err
	}

	if !info.IsDir() {
		return errors.New("output path is not a directory")
	}

	return nil
}

// GenerateOutputPath creates the full output path with the proper extension.
func GenerateOutputPath(outputDir, inputPath, randomName string) string {
	ext := filepath.Ext(inputPath)
	filename := randomName + ext

	return filepath.Join(outputDir, filename)
}
