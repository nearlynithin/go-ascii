package main

import (
	"fmt"
	"image"
	"log"

	"golang.org/x/term"
)


func getNewBounds(imageData image.Image) (int, int, *image.RGBA){
	bound := imageData.Bounds()
	imageSet:= image.NewRGBA(bound)

	terminalWidth, _,err := term.GetSize(0)
	if err!=nil{
		log.Println("Unable to get Terminal Size")
	}

	
	fmt.Printf("Terminal width := %d\n",terminalWidth)
	if terminalWidth>100{
		if bound.Dx()>1000{

			terminalWidth-=30
		}else{
			terminalWidth-=50
		}	
	}else{
		terminalWidth-=10
	}


	widthFactor := int(bound.Dx()/terminalWidth)
	heigthFactor:= widthFactor*2


	return widthFactor, heigthFactor, imageSet
}