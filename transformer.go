package main

import (
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

	widthFactor := int(bound.Dx()/terminalWidth+2)
	heigthFactor:= int((bound.Dy()*widthFactor)/bound.Dy()) +3

	return widthFactor, heigthFactor, imageSet
}