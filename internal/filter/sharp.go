package filter

import (
	"image"
	"image/color"
)

// ApplySharp applies the sharp filter to an image.
func ApplySharp(img image.Image, _ *Effect) image.Image {
	bounds := img.Bounds()
	output := image.NewRGBA(bounds)

	for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
		for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
			center := img.At(x, y)
			top := img.At(x, y-1)
			bottom := img.At(x, y+1)
			left := img.At(x-1, y)
			right := img.At(x+1, y)

			cr, cg, cb, _ := center.RGBA()
			tr, tg, tb, _ := top.RGBA()
			br, bg, bb, _ := bottom.RGBA()
			lr, lg, lb, _ := left.RGBA()
			rr, rg, rb, _ := right.RGBA()

			sharpR := int32(cr)*5 - int32(tr) - int32(br) - int32(lr) - int32(rr)
			sharpG := int32(cg)*5 - int32(tg) - int32(bg) - int32(lg) - int32(rg)
			sharpB := int32(cb)*5 - int32(tb) - int32(bb) - int32(lb) - int32(rb)

			if sharpR > RGBA<<8 {
				sharpR = RGBA << 8
			} else if sharpR < 0 {
				sharpR = 0
			}

			if sharpG > RGBA<<8 {
				sharpG = RGBA << 8
			} else if sharpG < 0 {
				sharpG = 0
			}

			if sharpB > RGBA<<8 {
				sharpB = RGBA << 8
			} else if sharpB < 0 {
				sharpB = 0
			}

			output.Set(x, y, color.RGBA{
				R: uint8(sharpR >> 8),
				G: uint8(sharpG >> 8),
				B: uint8(sharpB >> 8),
				A: Alpha,
			})
		}
	}

	return output
}
