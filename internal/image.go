package internal

import (
	"fmt"
	"image"
	"image/color"
	"os"
)

func toGrayScale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func avgPixel(image image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, image.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += toGrayScale(image.At(i, j))
			cnt++
		}
	}

	return sum / cnt
}

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func PrintAsciiImage(image image.Image) {
	ramp := "@%#*+=-:. "
	max := image.Bounds().Max
	scaleX, scaleY := 10, 5
	for y := 0; y < max.Y; y += scaleY {
		for x := 0; x < max.X; x += scaleX {
			avg := avgPixel(image, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp)*avg/65536]))
		}
		fmt.Println()
	}
}
