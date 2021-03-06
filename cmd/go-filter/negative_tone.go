package main

import (
	"image"
	"image/color"

	"github.com/urfave/cli/v2"
)

func NewNegativeSubCommand() *cli.Command {
	return &cli.Command{
		Name:   "negative",
		Usage:  "apply the negative filter",
		Action: applyNegativeFilter,
	}
}

func applyNegativeFilter(c *cli.Context) error {
	filePath := c.Args().First()
	e := new(Effect)

	file, err := e.ReadFile(filePath)
	if err != nil {
		return err
	}
	imgConf, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}
	width, height := imgConf.Width, imgConf.Height

	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	output := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			go func(x, y int) {
				r, g, b, _ := img.At(x, y).RGBA()
				sr, sg, sb := e.GetNegativeRGB(r/EightBits, g/EightBits, b/EightBits)
				filter := color.RGBA{
					R: sr,
					G: sg,
					B: sb,
					A: uint8(Alpha),
				}
				output.Set(x, y, filter)
			}(x, y)
		}
	}
	err = e.CreateFile(file, output, c.String("output"))
	if err != nil {
		return err
	}
	return nil

}
