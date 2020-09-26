package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	g := &game{
		mainChar: makeEtna(1, 2, 0),
		field: [][][]tile{
			[][]tile{
				[]tile{
					grassTile, grassTile, grassTile, grassTile,
				},
				[]tile{
					roadTile, roadTile, roadTile, roadTile,
				},
				[]tile{
					grassTile, snowTile, grassTile, grassTile,
				},
				[]tile{
					grassTile, grassTile, grassTile, grassTile,
				},
			},
			[][]tile{
				[]tile{
					nilTile, nilTile, grassTile, grassTile,
				},
				[]tile{
					nilTile, roadTile, roadTile, nilTile,
				},
				[]tile{
					nilTile, nilTile, nilTile, nilTile,
				},
				[]tile{
					nilTile, nilTile, nilTile, nilTile,
				},
			},
			[][]tile{
				[]tile{
					nilTile, nilTile, grassTile, grassTile,
				},
				[]tile{
					nilTile, nilTile, roadTile, nilTile,
				},
				[]tile{
					nilTile, nilTile, nilTile, nilTile,
				},
				[]tile{
					nilTile, nilTile, nilTile, nilTile,
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
