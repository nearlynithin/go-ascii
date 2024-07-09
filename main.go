package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)



func main(){
	test_img, err := os.Open("golang.png")
	if err!=nil{
		log.Fatal(err)
	}

	defer test_img.Close()
	
	imgData, _, err := image.Decode(test_img)
	
	bound := imgData.Bounds()
	imgSet := image.NewRGBA(bound)

	for y:=0; y<bound.Max.Y; y++ {
		for x:=0; x<bound.Max.X; x++{
			oldPixel := imgData.At(x,y)
			pixel:= color.GrayModel.Convert(oldPixel)
			imgSet.Set(x,y,pixel)
		}
	}

	outFile, err := os.Create("gray.png")
	if err!=nil{
		log.Fatal(err)
	}
	defer outFile.Close()

	 png.Encode(outFile, imgSet)
	
	
}