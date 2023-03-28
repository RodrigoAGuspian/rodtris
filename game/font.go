package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func NewFont() font.Face {
	typeText, err := opentype.Parse(fonts.PressStart2P_ttf)
	launchError(err)

	var generalFont font.Face

	generalFont, err = opentype.NewFace(typeText, &opentype.FaceOptions{
		Size: 24,
		DPI:  72,
	})

	launchError(err)

	return generalFont
}

func launchError(err error) {
	if err != nil {
		log.Fatal(err)

	}
}
