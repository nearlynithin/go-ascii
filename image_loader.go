package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

//retrieving data from the image
func loadImage(fileName string) (image.Image, error){
	filepath:= filepath.Join("assets", fileName)

	file, err := os.Open(filepath)
	if err!=nil{
		log.Println("File does not exist")
	}
	defer file.Close()

	imageData,_, err := image.Decode(file)
	if err!=nil{
		log.Println("Unable to read the file")
	}

	return imageData, nil
}