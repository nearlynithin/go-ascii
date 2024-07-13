package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/term"
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

	term_w, _, err:= term.GetSize(0)
	if err!=nil{
		log.Println("Unable to get terminal size")
	}

	term_w-=30

	factor_w:=int(bound.Dx()/term_w)
	factor_h:=int((bound.Dy()*factor_w)/bound.Dx())+4

	//  heightscale:=3
	//  widthscale:=2
	for y := bound.Min.Y; y < bound.Max.Y; y+=factor_h{
		for x := bound.Min.X; x < bound.Max.X; x+=factor_w{
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
