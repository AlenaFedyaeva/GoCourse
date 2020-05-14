package main

import (
	"GoCource/homework6/task2"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func ChessBoard() {
	width := 800
	height := 800

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	// cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, color.Black)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f,err := os.Create("image.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer f.Close()
	png.Encode(f, img)
}



func main() {
	task2.DrawLines()

	// ChessBoard()
	fmt.Println("\nThe end")
}
