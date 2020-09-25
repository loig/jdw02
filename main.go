package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	g := &game{
		field: [][][]tile{
			[][]tile{
				[]tile{
					grassTile, grassTile, grassTile, grassTile,
				},
				[]tile{
					roadTile, roadTile, roadTile, roadTile,
				},
				[]tile{
					grassTile, grassTile, grassTile, grassTile,
				},
				[]tile{
					grassTile, grassTile, grassTile, grassTile,
				},
			},
		},
	}
	ebiten.SetWindowSize(1280, 640)
	ebiten.SetWindowTitle("Jdw02")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
