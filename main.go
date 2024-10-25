package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	// Open a video capture device
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", err)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Hello")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error reading from webcam\n")
			return
		}
		if img.Empty() {
			continue
		}

		width := img.Cols()
		height := img.Rows()
		centerX := width / 2
		centerY := height / 2

		r := img.GetUCharAt(centerY, centerX*3+2)
		g := img.GetUCharAt(centerY, centerX*3+1)
		b := img.GetUCharAt(centerY, centerX*3)

		fmt.Printf("Coordinates: (%d, %d) - RGB: (%d, %d, %d)\n", centerX, centerY, r, g, b)

		window.IMShow(img)

		if window.WaitKey(1) == 27 {
			break
		}
	}
}
