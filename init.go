package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func init() {

	grassTileImage, _, error := ebitenutil.NewImageFromFile("assets/grass.png", ebiten.FilterDefault)
	if error != nil {
		log.Fatal(error)
	}
	grassTile = tile{image: grassTileImage}

	roadTileImage, _, error := ebitenutil.NewImageFromFile("assets/road.png", ebiten.FilterDefault)
	if error != nil {
		log.Fatal(error)
	}
	roadTile = tile{image: roadTileImage}
}
