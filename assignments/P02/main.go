package main

import (
	"P02/imagemod/imageManipulator"
	"fmt"
)

func main() {
	im := imageManipulator.NewImageManipulator(800, 600)
	im2, err := imageManipulator.NewImageManipulatorWithImage("mustangs.jpg")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	im.DrawRectangle(420, 69, 42, 42)
	im2.DrawRectangle(200, 100, 500, 350)

	im.SaveToFile("silly-little-image.png")
	im2.SaveToFile("mustangs.png")
}
