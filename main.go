package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
)

func main() {
	test_img, err := os.Open("assets/golang.png")
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

	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal("error creating file")
	}
	defer f.Close()

	heightscale:=2
	for y := bound.Min.Y; y < bound.Max.Y; y+=heightscale {
		for x := bound.Min.X; x < bound.Max.X; x++{
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
				f.WriteString("@")
			case intensity>=191:
				f.WriteString("%")
			case intensity>=127:
				f.WriteString("#")
			case intensity>=64:
				f.WriteString("$")
			default:
				f.WriteString("   ")
			}
		}
		f.WriteString("\n")
	}
}
