package main

import (
	"github.com/hajimehoshi/ebiten"
)

type game struct {
	field    [][][]tile
	mainChar character
}

func (g *game) Update(screen *ebiten.Image) error {

	g.mainChar.updateAnim()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	yshift := float64(1600 - 150)
	xshift := float64(3200 - 500)

	for y := 0; y < len(g.field[0]); y++ {
		for x := len(g.field[0][0]) - 1; x >= 0; x-- {
			for z := range g.field {
				if g.field[z][y][x].image != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*256+y*256)+xshift, float64(y*154-x*144-z*120)+yshift)

					screen.DrawImage(g.field[z][y][x].image, op)
				}
				if g.mainChar.posZ == z {
					if g.mainChar.posX >= float64(x)-0.5 &&
						g.mainChar.posX < float64(x)+0.5 {
						if g.mainChar.posY >= float64(y)-0.5 &&
							g.mainChar.posY < float64(y)+0.5 {
							op := &ebiten.DrawImageOptions{}
							op.GeoM.Scale(3, 3)
							op.GeoM.Translate(float64(x*256+y*256)+xshift+50, float64(y*154-x*144-z*120)+yshift-370)

							screen.DrawImage(g.mainChar.currentImage, op)
						}
					}
				}
			}
		}
	}

	return
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 6400, 3200
}
