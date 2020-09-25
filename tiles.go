package main

import "github.com/hajimehoshi/ebiten"

type tile struct {
	image *ebiten.Image
}

var grassTile tile
var roadTile tile
