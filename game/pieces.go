package game

import (
	_ "embed"
	_ "image/png"
)

//go:embed images/bloque.png
var imgBlock1 []byte

//go:embed images/bloqueA.png
var imgBlock2 []byte

//go:embed images/bloqueAz.png
var imgBlock3 []byte

//go:embed images/bloqueNa.png
var imgBlock4 []byte

//go:embed images/bloqueP.png
var imgBlock5 []byte

//go:embed images/bloqueR.png
var imgBlock6 []byte

//go:embed images/bloqueVe.png
var imgBlock7 []byte

//go:embed images/bloqueV.png
var imgBlock8 []byte

//go:embed images/bloqueTu.png
var imgBlock9 []byte

//go:embed images/bg.png
var imgBlockBG []byte

var allPieces = []*Piece{
	NewPiece([][]int{
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	}, imgBlock1),
	NewPiece([][]int{
		{0, 0, 0, 1},
		{0, 1, 1, 1},
	}, imgBlock2),
	NewPiece([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 0},
	}, imgBlock3),
	NewPiece([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
	}, imgBlock4),
	NewPiece([][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
	}, imgBlock5),
	NewPiece([][]int{
		{0, 1, 1, 0},
		{1, 1, 0, 0},
	}, imgBlock6),
	NewPiece([][]int{
		{1, 1, 0, 0},
		{0, 1, 1, 0},
	}, imgBlock7),
	NewPiece([][]int{
		{0, 1, 0, 0},
		{0, 1, 1, 0},
	}, imgBlock8),
	NewPiece([][]int{
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}, imgBlock9),
}
