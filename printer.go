package main

import (
	"fmt"
	"image/color"
)


func getRGB(newPixel color.Color) (uint8,uint8,uint8){
	red, green, blue, _ := newPixel.RGBA()
	red8 := uint8(red>>8)
	green8 := uint8(green>>8)
	blue8 := uint8(blue>>8)

		return red8, green8, blue8
}

func getIntensity(red, green, blue uint8) float64 {
	return (0.2126*float64(red) + 0.7152*float64(green) + 0.0722*float64(blue))
}


func asciiPrinter(intensity float64){
	switch {
	case intensity>= 240:
		fmt.Print("@")
	case intensity>=200:
		fmt.Print("%")
	case intensity>=160:
		fmt.Print("#")
	case intensity>=120:
		fmt.Print("$")
	case intensity>=80:
		fmt.Print("*")
	case intensity>=40:
		fmt.Print("`")
	default:
		fmt.Print(" ")
	}
}