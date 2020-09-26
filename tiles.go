package main

import "github.com/hajimehoshi/ebiten"

type tile struct {
	image *ebiten.Image
}

var nilTile tile = tile{image: nil}
var grassTile tile
var roadTile tile
