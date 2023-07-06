package main

import (
	"fmt"
	"image"
	"time"

	"github.com/muesli/termenv"
	"github.com/nfnt/resize"
	"gocv.io/x/gocv"
)

var asciiChars = "MND8OZ$7I?+=~:."

func mapGrayToASCII(grayValue uint32) byte {
	pos := int((float32(grayValue) / 65535.0) * float32(len(asciiChars)-1))
	return asciiChars[pos]
}

func convertFrameToASCII(img image.Image, width uint, usecolor bool) string {
	newHeight := uint(float64(img.Bounds().Dy()) * (float64(width) / float64(img.Bounds().Dx())) / 2)
	img = resize.Resize(width, newHeight, img, resize.Lanczos3)

	var ascii string
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b)
			asciiChar := mapGrayToASCII(uint32(gray))
			if usecolor {
				color := termenv.ColorProfile().Color(fmt.Sprintf("#%02x%02x%02x", uint8(r>>8), uint8(g>>8), uint8(b>>8)))
				ascii += termenv.String(string(asciiChar)).Foreground(color).String()
			} else {
				ascii += string(asciiChar)
			}
		}
		ascii += "\n"
	}
	return ascii
}

func main() {
	webcam, err := gocv.OpenVideoCapture(0) // 0 is the id of the first webcam on your system
	if err != nil {
		fmt.Println("Error: cannot open video capture")
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Cannot read device")
			return
		}
		if img.Empty() {
			continue
		}

		goimg, _ := img.ToImage()
		ascii := convertFrameToASCII(goimg, 80, false) // You can replace 80 with your desired width
		fmt.Print("\033[H\033[2J", ascii)

		time.Sleep(time.Millisecond * 33) // 33 milliseconds is approximately 30 frames per second
	}
}
