package game

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X int
	Y int
}

func (current *Position) Add(new Position) {
	current.X += new.X
	current.Y += new.Y
}

func createImage(imgData []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
