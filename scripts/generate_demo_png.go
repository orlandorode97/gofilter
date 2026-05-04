// Command generate_demo_png writes fixtures/demo.png for demos (e.g. VHS recordings).
// Run from the repository root: go run ./scripts/generate_demo_png.go
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fixturesDir := filepath.Join(wd, "fixtures")
	if err := os.MkdirAll(fixturesDir, 0o755); err != nil {
		panic(err)
	}

	outPath := filepath.Join(fixturesDir, "demo.png")

	img := image.NewRGBA(image.Rect(0, 0, 240, 180))
	for y := 0; y < 180; y++ {
		for x := 0; x < 240; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8((x + y) % 256),
				G: uint8(x % 256),
				B: uint8(y % 256),
				A: 255,
			})
		}
	}

	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}
}
