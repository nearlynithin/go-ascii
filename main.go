package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"strconv"
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

	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			oldPixel := imgData.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			imgSet.Set(x, y, pixel)
			c := imgSet.At(x, y)
			r, g, b, a := c.RGBA()
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)
			_, err := f.WriteString(strconv.Itoa(int(r8)) + " " + strconv.Itoa(int(g8)) + " " + strconv.Itoa(int(b8)) + " " + strconv.Itoa(int(a8)) + "\n")
			if err != nil {
				log.Printf("Error writing RGB values: %v", err)
			}
		}
	}
}
