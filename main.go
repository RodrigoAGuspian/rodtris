package main

import (
	"log"
	"main/game"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const NAME = "Rodtris"

func main() {
	rand.Seed(time.Now().UnixNano())
	newGame := game.NewGame()

	ebiten.SetWindowSize(game.SCREEN_WIDTH, game.SCREEN_HEIGHT)
	ebiten.SetWindowTitle(NAME)

	if err := ebiten.RunGame(newGame); err != nil {
		log.Fatal(err)
	}
}
