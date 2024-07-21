package main

import (
	"flag"
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

func main() {
	// Supported input formats : jpg, jpeg, png
	//Get the file name
	// Define the flag for color printing
	useColor := flag.Bool("color", false, "Print images in color")
	flag.Parse()

	// Get the remaining non-flag arguments (positional arguments)
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Please provide an image file name as the argument.")
	}
	fileName := args[0]

	// Load the image data (pixels of image type)
	imageData, err := loadImage(fileName)
	if err != nil {
		log.Fatal(err)
	}


	//Number of pixels on the X and Y axis
	bound:= imageData.Bounds()

	//Creating a new Image type with reduced width and heigth to fit the terminal
	widthFactor, heigthFactor, _ := getNewBounds(imageData) //ImageSet is the third return
	// fmt.Printf("Dimensions X:= %d and Y:= %d\n wf := %d hf:= %d\n",bound.Dx(),bound.Dy(), widthFactor, heigthFactor)

	//Printing in color
	
	if *useColor{
			//Printing each pixel of the new image while converting each oldpixel to grayscale and mapping a suitable ascii character according to luminosity of the pixel
			for y := bound.Min.Y; y < bound.Max.Y; y+=heigthFactor{
				for x := bound.Min.X; x < bound.Max.X; x+=widthFactor{
					oldPixel := imageData.At(x, y)
					pixel := color.GrayModel.Convert(oldPixel)
					// imageSet.Set(x, y, pixel)
					// newPixel := imageSet.At(x, y)
					r, g, b := getRGB(pixel)
					cr,cg,cb:=getRGB(oldPixel)
					intensity := getIntensity(r,g,b)
					asciiColorPrinter(intensity,cr,cg,cb)
				}
				fmt.Print("\n")
			}
	}else{
		for y := bound.Min.Y; y < bound.Max.Y; y+=heigthFactor{
			for x := bound.Min.X; x < bound.Max.X; x+=widthFactor{
				oldPixel := imageData.At(x, y)
				pixel := color.GrayModel.Convert(oldPixel)
				// imageSet.Set(x, y, pixel)
				// newPixel := imageSet.At(x, y)
				r, g, b := getRGB(pixel)
				intensity := getIntensity(r,g,b)
				asciiPrinter(intensity)
			}
			fmt.Print("\n")
		}
	}


}


