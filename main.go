package main

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// Supported input formats : jpg, jpeg, png
	//Get the file name
	fileName := os.Args[1]

	//Load the image data (pixels of image type)
	imageData,_ := loadImage(fileName)

	//Number of pixels on the X and Y axis
	bound:= imageData.Bounds()

	//Creating a new Image type with reduced width and heigth to fit the terminal
	widthFactor, heigthFactor, imageSet := getNewBounds(imageData)
	// fmt.Printf("Dimensions X:= %d and Y:= %d\n wf := %d hf:= %d\n",bound.Dx(),bound.Dy(), widthFactor, heigthFactor)

	//Printing each pixel of the new image while converting each oldpixel to grayscale and mapping a suitable ascii character according to luminosity of the pixel
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
