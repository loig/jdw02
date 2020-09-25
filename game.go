package main

import (
	"github.com/hajimehoshi/ebiten"
)

type game struct {
	field [][][]tile
}

func (g *game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	yshift := float64(1600 - 150)
	xshift := float64(3200 - 500)

	for _, level := range g.field {
		for y, line := range level {
			for x := len(line) - 1; x >= 0; x-- {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*256+y*256)+xshift, float64(y*154-x*144)+yshift)

				screen.DrawImage(line[x].image, op)
			}
		}
	}

	return
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 6400, 3200
}
