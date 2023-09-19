package main

import imageManipulator "imagemod/imageManipulator"

func main() {
	im := imageManipulator.NewImageManipulator(800, 600)

	im.DrawRectangle(420, 69, 42, 42)

	im.SaveToFile("silly-little-image.png")
}
