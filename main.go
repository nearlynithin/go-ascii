package main

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {

	fileName := os.Args[1]
	imageData,_ := loadImage(fileName)
	bound:= imageData.Bounds()

	widthFactor, heigthFactor, imageSet := getNewBounds(imageData)


	for y := bound.Min.Y; y < bound.Max.Y; y+=heigthFactor{
		for x := bound.Min.X; x < bound.Max.X; x+=widthFactor{
			oldPixel := imageData.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			imageSet.Set(x, y, pixel)
			newPixel := imageSet.At(x, y)
			r, g, b := getRGB(newPixel)
			intensity := getIntensity(r,g,b)
			asciiPrinter(intensity)
		}
		fmt.Print("\n")
	}
}
