package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) transferPieceToGameZone() {
	piece := g.currentPiece
	piecePosition := g.piecePosition

	for dy, row := range piece.Blocks {
		for dx, value := range row {
			if value != BLOCK {
				continue
			}

			gameZonePosition := &Position{X: piecePosition.X + dx, Y: piecePosition.Y + dy}

			g.gameZone[gameZonePosition.Y][gameZonePosition.X] = piece.Image
		}
	}
}

func (g *Game) insideGameZone(deltaPosition Position) bool {
	piecePosition := *g.piecePosition
	piecePosition.Add(deltaPosition)
	return g.pieceInsideGameZone(g.currentPiece, piecePosition)
}

func (g *Game) pieceInsideGameZone(piece *Piece, piecePosition Position) bool {
	for dy, row := range piece.Blocks {
		for dx, value := range row {
			if value == BLOCK {
				screenPosition := &Position{X: piecePosition.X + dx, Y: piecePosition.Y + dy}

				if screenPosition.X < 0 || screenPosition.X >= int(g.gameZoneSize.Width) {
					return false
				}

				if screenPosition.Y < 0 || screenPosition.Y >= int(g.gameZoneSize.Height) {
					return false
				}

				if g.gameZone[screenPosition.Y][screenPosition.X] != nil {
					return false
				}

			}
		}
	}

	return true
}

func (g *Game) fetchNextPiece() {
	getNextPiece := func() *Piece {
		return g.pieces[rand.Intn(len(g.pieces))]
	}

	if g.nextPiece == nil {
		g.nextPiece = getNextPiece()
	}

	g.currentPiece = g.nextPiece
	g.piecePosition = &Position{X: int(g.gameZoneSize.Width)/2 - 1, Y: 0}
	g.nextPiece = getNextPiece()

}

func (g *Game) checkForLines() int {
	lines := []int{}

	for y, row := range g.gameZone {
		full := true
		for _, cellImage := range row {
			if cellImage == nil {
				full = false
				break
			}
		}
		if full {
			lines = append(lines, y)
		}
	}

	for _, y := range lines {
		emptyRow := [][]*ebiten.Image{make([]*ebiten.Image, g.gameZoneSize.Width)}
		g.gameZone = append(append(emptyRow, g.gameZone[0:y]...), g.gameZone[(y+1):]...)
	}

	return len(lines)
}

func (g *Game) rotatePiece() *Piece {
	newPiece := *g.currentPiece
	newPiece.Blocks = make([][]int, len(g.currentPiece.Blocks[0]))

	for y := range newPiece.Blocks {
		newPiece.Blocks[y] = make([]int, len(g.currentPiece.Blocks))
	}

	for y, row := range g.currentPiece.Blocks {
		for x := range row {
			newPiece.Blocks[x][len(g.currentPiece.Blocks)-1-y] = g.currentPiece.Blocks[y][x]
		}
	}

	return &newPiece
}
