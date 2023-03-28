package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	SCREEN_WIDTH  = 480
	SCREEN_HEIGHT = 480
	SCORE         = "SCORE"
	GAME_OVER     = "GAME OVER!"
	PRESS_SPACE   = "PRESS SPACE"
	TO_PLAY       = "TO PLAY"
	NEXT          = "NEXT"
)
const (
	GameStateGamerOver GameState = iota
	GameStatePlaying
)

type Size struct {
	Width  uint
	Height uint
}

type GameState int

type Game struct {
	dropTicks   uint
	elapsedDrop uint

	score       int
	state       GameState
	attractMode bool
	pieces      []*Piece
	txtFont     font.Face

	nextPiece     *Piece
	currentPiece  *Piece
	piecePosition *Position

	gameZoneSize Size
	gameZone     [][]*ebiten.Image
	bgBlockImage *ebiten.Image

	input            Input
	inputAttractMode Input
	inputKeyboard    Input
}

func (g *Game) Start() {
	g.state = GameStatePlaying
	g.score = 0
	g.attractMode = true
	g.input = g.inputAttractMode
	g.elapsedDrop = 0

	g.gameZone = make([][]*ebiten.Image, g.gameZoneSize.Height)

	for y := range g.gameZone {
		g.gameZone[y] = make([]*ebiten.Image, g.gameZoneSize.Width)
	}

	g.fetchNextPiece()
}

func (g *Game) StartPlay() {
	g.Start()
	g.attractMode = false
	g.input = g.inputKeyboard
}

func (g *Game) Update() error {
	g.elapsedDrop += 1

	switch g.state {
	case GameStatePlaying:
		if g.elapsedDrop > g.dropTicks {
			g.processInput(DOWN)
			g.elapsedDrop = 0
			return nil
		}

		key := g.input.Read()
		if key != nil {
			g.processInput(*key)
		}

		if g.attractMode && g.inputKeyboard.IsSpacePressed() {
			g.StartPlay()
		}

	case GameStateGamerOver:
		if g.input.IsSpacePressed() {
			if g.attractMode {
				g.Start()
			} else {
				g.StartPlay()
			}
		}
	}

	return nil
}

func (g *Game) processPiece() bool {
	g.transferPieceToGameZone()
	linesRemoves := g.checkForLines()
	g.updateScore(linesRemoves)
	g.fetchNextPiece()

	stopProcess := false
	deltaPosition := Position{}

	if !g.insideGameZone(deltaPosition) {
		g.state = GameStateGamerOver
		stopProcess = true
	}
	return stopProcess
}

func (g *Game) processInput(key ebiten.Key) {
	switch key {
	case ebiten.KeyDown:
		deltaPosition := Position{X: 0, Y: 1}
		if g.insideGameZone(deltaPosition) {
			g.piecePosition.Add(deltaPosition)
		} else {
			stopProcess := g.processPiece()
			if stopProcess {
				return
			}
		}
	case ebiten.KeyLeft:
		deltaPosition := Position{X: -1, Y: 0}
		if g.insideGameZone(deltaPosition) {
			g.piecePosition.Add(deltaPosition)
		}

	case ebiten.KeyRight:
		deltaPosition := Position{X: 1, Y: 0}
		if g.insideGameZone(deltaPosition) {
			g.piecePosition.Add(deltaPosition)
		}

	case ebiten.KeyUp:
		newPiece := g.rotatePiece()
		if g.pieceInsideGameZone(newPiece, *g.piecePosition) {
			g.currentPiece = newPiece
		}

	}

}

func (g *Game) drawText(screen *ebiten.Image, gameZonePosition *Position) {
	boardBlockWidth, _ := g.bgBlockImage.Size()
	boardWidth := int(g.gameZoneSize.Width) * boardBlockWidth
	text.Draw(screen, SCORE, g.txtFont, boardWidth+gameZonePosition.X*2, gameZonePosition.Y*2, color.White)
	text.Draw(screen, fmt.Sprintf("%08d", g.score), g.txtFont, boardWidth+gameZonePosition.X*2, gameZonePosition.Y*2+8, color.White)

	if g.state == GameStateGamerOver {
		dy := 148
		text.Draw(screen, GAME_OVER, g.txtFont, boardWidth+gameZonePosition.X*2, gameZonePosition.Y*2+dy, color.White)
		text.Draw(screen, PRESS_SPACE, g.txtFont, boardWidth+gameZonePosition.X*2, gameZonePosition.Y*2+8+dy, color.White)
	}

	if g.attractMode {
		dy := 148
		text.Draw(screen, PRESS_SPACE, g.txtFont, boardWidth+gameZonePosition.X*2, gameZonePosition.Y*2+dy, color.White)
		text.Draw(screen, TO_PLAY, g.txtFont, boardWidth+gameZonePosition.X*2, gameZonePosition.Y*2+8+dy, color.White)
	}

	dy := 48
	text.Draw(screen, NEXT, g.txtFont, boardWidth+gameZonePosition.X*2, gameZonePosition.Y*2+dy, color.White)
}

func (g *Game) updateScore(lines int) {
	perLineScore := 100
	g.score += lines * perLineScore
	if lines > 1 {
		bonus := perLineScore / 2
		for i := 0; i < int(lines); i++ {
			g.score += bonus
			bonus *= 2
		}
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	gameZonePosition := &Position{X: 16, Y: 16}
	g.drawText(screen, gameZonePosition)

	gameZone := g.gameZone

	for y, row := range gameZone {
		for x, cellImage := range row {
			if cellImage == nil {
				cellImage = g.bgBlockImage
			}

			w, h := cellImage.Size()
			screenPosition := &Position{
				X: gameZonePosition.X + x*w,
				Y: gameZonePosition.Y + y*h,
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(screenPosition.X), float64(screenPosition.Y))
			screen.DrawImage(cellImage, op)
		}
	}

	if g.currentPiece != nil {
		g.currentPiece.Draw(screen, gameZonePosition, g.piecePosition)
	}

	if g.nextPiece != nil {
		nextPosition := &Position{X: int(math.Round(SCREEN_WIDTH * 0.25)), Y: int(math.Round(SCREEN_HEIGHT * 0.22))}
		g.nextPiece.Draw(screen, nextPosition, &Position{})
	}

}

func (g *Game) Layout(outsideWidth, outsideHeigth int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func NewGame() *Game {
	ebiten.SetMaxTPS(18)

	game := &Game{
		txtFont:          NewFont(),
		inputAttractMode: NewAttractModeInput(),
		inputKeyboard:    &KeyboardInput{},
		dropTicks:        4,
		pieces:           allPieces,
		gameZoneSize:     Size{Width: 10, Height: 24},
		bgBlockImage:     createImage(imgBlockBG),
	}

	game.Start()

	return game
}
