package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const BLOCK = 1

type Piece struct {
	Blocks [][]int
	Image  *ebiten.Image
}

func NewPiece(blocks [][]int, imgData []byte) *Piece {
	return &Piece{
		Blocks: blocks, Image: createImage(imgData),
	}
}

func (p *Piece) Draw(screen *ebiten.Image, gameZonePosition *Position, piecePos *Position) {
	w, h := p.Image.Size()

	for dy, row := range p.Blocks {
		for dx, value := range row {
			if value == BLOCK {
				screenPosition := &Position{
					X: gameZonePosition.X + (piecePos.X+dx)*w,
					Y: gameZonePosition.Y + (piecePos.Y+dy)*h,
				}

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(screenPosition.X), float64(screenPosition.Y))
				screen.DrawImage(p.Image, op)
			}
		}
	}
}
