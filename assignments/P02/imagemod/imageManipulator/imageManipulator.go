package imageManipulator

import (
	"github.com/fogleman/gg"
)

type ImageManipulator struct {
	Image *gg.Context
}

func NewImageManipulator(width int, height int) *ImageManipulator {
	img := gg.NewContext(width, height)
	return &ImageManipulator{Image: img}
}

func (im *ImageManipulator) SaveToFile(filename string) error {
	return im.Image.SavePNG(filename)
}

func (im *ImageManipulator) DrawRectangle(x float64, y float64, width float64, height float64) {
	im.Image.DrawRectangle(x, y, width, height)
	im.Image.Stroke()
}
