package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	test_img, err := os.Open("assets/pfp2.png")
	if err != nil {
		log.Fatal(err)
	}
	defer test_img.Close()
	
	imgData, _, err := image.Decode(test_img)
	if err != nil {
		log.Fatal(err)
	}

	bound := imgData.Bounds()
	imgSet := image.NewRGBA(bound)

	 heightscale:=3
	 widthscale:=2
	 factor:=6
	for y := bound.Min.Y; y < bound.Max.Y; y+=heightscale*factor{
		for x := bound.Min.X; x < bound.Max.X; x+=widthscale*factor{
			oldPixel := imgData.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			imgSet.Set(x, y, pixel)
			c := imgSet.At(x, y)
			r, g, b, _ := c.RGBA()
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			// a8 := uint8(a >> 8)
			
			intensity := (r8+g8+b8/3)
			switch {
			case intensity>= 255:
				fmt.Print("@")
			case intensity>=191:
				fmt.Print("%")
			case intensity>=127:
				fmt.Print("#")
			case intensity>=64:
				fmt.Print("$")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
