package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/urfave/cli/v2"
)

const (
	RedWaveLength   float64 = 0.21
	GreenWaveLength float64 = 0.72
	BlueWaveLength  float64 = 0.07
	RGBA            int     = 255
)

func NewGraySubCommand() *cli.Command {
	return &cli.Command{
		Name:   "gray",
		Usage:  "apply the gray filter",
		Action: applyGrayFilter,
	}
}

func applyGrayFilter(c *cli.Context) error {
	filePath := c.Args().First()

	file, err := effect.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	imgConf, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}
	width, height := imgConf.Width, imgConf.Height

	// reset io.Reader
	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	//new image to create
	output := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			go func(x, y int) {
				r, g, b, _ := img.At(x, y).RGBA()
				rgb := effect.GetGrayRGB(r/EightBits, g/EightBits, b/EightBits)
				filter := color.RGBA{
					R: rgb,
					G: rgb,
					B: rgb,
					A: uint8(Alpha),
				}
				output.Set(x, y, filter)
			}(x, y)
		}
	}
	err = effect.CreateFile(file, output, c.String("output"))
	if err != nil {
		return err
	}
	return nil
}
