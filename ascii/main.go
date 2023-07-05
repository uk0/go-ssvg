package main

import (
	"fmt"
	"github.com/muesli/termenv"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

var asciiChars = "MND8OZ$7I?+=~:,.."

func mapGrayToASCII(grayValue uint32) byte {
	pos := int((float32(grayValue) / 65535.0) * float32(len(asciiChars)-1))
	return asciiChars[pos]
}

func convertToColorASCII(img image.Image, width uint) {
	newHeight := uint(float64(img.Bounds().Dy()) * (float64(width) / float64(img.Bounds().Dx())) / 2)
	img = resize.Resize(width, newHeight, img, resize.Lanczos3)

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b)
			asciiChar := mapGrayToASCII(uint32(gray))

			color := termenv.ColorProfile().Color(fmt.Sprintf("#%02x%02x%02x", uint8(r>>8), uint8(g>>8), uint8(b>>8)))
			fmt.Print(termenv.String(string(asciiChar)).Foreground(color).String())
		}
		fmt.Println()
	}
}

func convertToASCII(img image.Image, width uint) {
	newHeight := uint(float64(img.Bounds().Dy()) * (float64(width) / float64(img.Bounds().Dx())) / 2)
	img = resize.Resize(width, newHeight, img, resize.Lanczos3)

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b)
			asciiChar := mapGrayToASCII(uint32(gray))
			fmt.Print(string(asciiChar))
		}
		fmt.Println()
	}
}

func main() {
	imgFile, err := os.Open("/Users/firshme/Desktop/work/go-ssvg/ascii/misc-ascii-mona.jpg") // Replace with your image file
	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	convertToASCII(img, 120)      // You can replace 80 with your desired width
	convertToColorASCII(img, 120) // You can replace 80 with your desired width
}
