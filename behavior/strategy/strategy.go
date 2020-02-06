package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

var (
	output string
)

func init() {
	flag.StringVar(&output, "output", "console", "The output to use between 'console' and 'image' file")
}

type PrintStrategy interface {
	Draw() error
}

type ConsoleSquare struct{}

func (c *ConsoleSquare) Draw() error {
	fmt.Println("Square")
	return nil
}

type ImageSquare struct {
	DestinationFilePath string
}

func (t *ImageSquare) Draw() error {
	width := 800
	height := 600

	origin := image.Point{0, 0}

	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}

	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 110, G: 65, B: 186, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})

	squareImg := image.NewRGBA(square)

	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("Error opening image")
	}
	defer w.Close()

	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}

	return nil
}

func main() {
	flag.Parse()

	var activeStrategy PrintStrategy

	switch output {
	case "console":
		activeStrategy = &ConsoleSquare{}
	case "image":
		activeStrategy = &ImageSquare{"./image.jpg"}
	default:
		activeStrategy = &ConsoleSquare{}
	}

	err := activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
