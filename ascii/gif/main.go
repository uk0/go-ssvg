package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"time"

	"github.com/muesli/termenv"
	"github.com/nfnt/resize"
)

var asciiChars = "MND8OZ$7I?+=~:."

func mapGrayToASCII(grayValue uint32) byte {
	pos := int((float32(grayValue) / 65535.0) * float32(len(asciiChars)-1))
	return asciiChars[pos]
}

func convertFrameToASCII(img image.Image, width uint) string {
	newHeight := uint(float64(img.Bounds().Dy()) * (float64(width) / float64(img.Bounds().Dx())) / 2)
	img = resize.Resize(width, newHeight, img, resize.Lanczos3)

	var ascii string
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b)
			asciiChar := mapGrayToASCII(uint32(gray))

			color := termenv.ColorProfile().Color(fmt.Sprintf("#%02x%02x%02x", uint8(r>>8), uint8(g>>8), uint8(b>>8)))
			ascii += termenv.String(string(asciiChar)).Foreground(color).String()
		}
		ascii += "\n"
	}
	return ascii
}

func main() {
	gifFile, err := os.Open("/Users/firshme/Desktop/work/go-ssvg/ascii/gif/test1.gif") // Replace with your gif file
	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}
	defer gifFile.Close()

	gifImage, err := gif.DecodeAll(gifFile)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	for _, frame := range gifImage.Image {
		img := image.NewPaletted(frame.Bounds(), palette.Plan9)
		draw.Draw(img, frame.Bounds(), frame, image.ZP, draw.Src)

		ascii := convertFrameToASCII(img, 80) // You can replace 80 with your desired width
		fmt.Print("\033[H\033[2J", ascii)

		time.Sleep(time.Duration(gifImage.Delay[0]*10) * time.Millisecond)
	}
}
