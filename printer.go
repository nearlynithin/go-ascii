package main

import (
	"fmt"
	"image/color"

	stdcolor "github.com/gookit/color"
)

//handler to get 8 bit RGB values of a pixel
func getRGB(newPixel color.Color) (uint8,uint8,uint8){
	red, green, blue, _ := newPixel.RGBA()
	red8 := uint8(red>>8)
	green8 := uint8(green>>8)
	blue8 := uint8(blue>>8)

		return red8, green8, blue8
}


// The lightness method averages the most prominent and least prominent colors: (max(R, G, B) + min(R, G, B)) / 2.

// The average method simply averages the values: (R + G + B) / 3.

// The formula for luminosity is 0.21 R + 0.72 G + 0.07 B.
// calculating intensity of a pixel based on the luminosity (Better results)
func getIntensity(red uint8, green uint8, blue uint8) float64 {
	return (0.2126*float64(red) + 0.7152*float64(green) + 0.0722*float64(blue))
}


//Priting ascii characters based on the intensity
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

func asciiColorPrinter(intensity float64, c_red uint8, c_green uint8, c_blue uint8){
	switch {
	case intensity>= 240:
		stdcolor.RGB(c_red,c_green,c_blue).Print("@")
	case intensity>=200:
		// fmt.Print("%")
		stdcolor.RGB(c_red,c_green,c_blue).Print("%")
	case intensity>=160:
		// fmt.Print("#")
		stdcolor.RGB(c_red,c_green,c_blue).Print("#")
	case intensity>=120:
		// fmt.Print("$")
		stdcolor.RGB(c_red,c_green,c_blue).Print("$")
	case intensity>=80:
		// fmt.Print("*")
		stdcolor.RGB(c_red,c_green,c_blue).Print("*")
	case intensity>=40:
		// fmt.Print("`")
		stdcolor.RGB(c_red,c_green,c_blue).Print("`")
	default:
		// fmt.Print(" ")
		stdcolor.RGB(c_red,c_green,c_blue).Print(" ")
	}
}