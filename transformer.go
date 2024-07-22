package main

import (
	"image"
	"log"
	"os"

	"golang.org/x/term"
)

//Creating a new blank image.
func getNewBounds(imageData image.Image) (int, int, *image.RGBA){
	bound := imageData.Bounds()
	imageSet:= image.NewRGBA(bound)

	//os.Stdout.Fd() for non-Unix systems.
	//on UNIX-like systems its 0
	terminalWidth, _,err := term.GetSize(int(os.Stdout.Fd()))
	if err!=nil{
		log.Println("Unable to get Terminal Size")
	}

	//TODO : Better algorithm for calculating aspect ratio as per the terminal size
	// fmt.Printf("Terminal width := %d\n",terminalWidth)

	//Remove the following code block to get rid of the blank space at the right side of the terminal
	//From here
	if terminalWidth>100{
		if bound.Dx()>1000{

			terminalWidth-=30
		}else{
			terminalWidth-=50
		}	
	}else{
		terminalWidth-=10
	}
	//Till here

	//basically the steps size for the for loops
	widthFactor := int(bound.Dx()/terminalWidth)
	heigthFactor:= widthFactor*2


	return widthFactor, heigthFactor, imageSet
}