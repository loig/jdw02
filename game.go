package main

import (
	"github.com/hajimehoshi/ebiten"
)

type game struct {
	field    [][][]tile
	mainChar character
}

func (g *game) Update(screen *ebiten.Image) error {

	var hasMoved bool

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.mainChar.direction = 0
		hasMoved = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.mainChar.direction = 1
		hasMoved = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.mainChar.direction = 2
		hasMoved = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.mainChar.direction = 3
		hasMoved = true
	}

	if hasMoved {
		g.mainChar.moveType = walking
		g.moveChar()
	} else {
		g.mainChar.moveType = idle
	}

	g.mainChar.updateAnim()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	yshift := float64(1600 - 150)
	xshift := float64(3200 - 500)

	maxX := len(g.field[0][0])
	maxY := len(g.field[0])
	maxZ := len(g.field)

	displayCharNext := false
	displayedChar := false
	for startX := maxX - 1; startX > -maxY; startX-- {
		isCharLine := false
		for z := 0; z < maxZ; z++ {
			y := 0
			for x := startX; x < maxX; x++ {
				if x >= 0 && y < maxY {
					if g.field[z][y][x].image != nil {
						op := &ebiten.DrawImageOptions{}
						op.GeoM.Translate(float64(x*256+y*256)+xshift, float64(y*154-x*144-z*120)+yshift)
						screen.DrawImage(g.field[z][y][x].image, op)
					}
					if g.mainChar.posZ == z &&
						g.mainChar.posX >= float64(x)-0.5 &&
						g.mainChar.posX <= float64(x)+0.5 &&
						g.mainChar.posY >= float64(y)-0.5 &&
						g.mainChar.posY <= float64(y)+0.5 {
						isCharLine = true
					}
				}
				y++
			}
			if displayCharNext && !displayedChar &&
				z == g.mainChar.posZ {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(3, 3)
				op.GeoM.Translate(g.mainChar.posX*256+g.mainChar.posY*256+xshift+50, g.mainChar.posY*154-g.mainChar.posX*144-float64(g.mainChar.posZ*120)+yshift-370)

				screen.DrawImage(g.mainChar.currentImage, op)
				displayedChar = true
			}
		}
		if isCharLine {
			displayCharNext = true
		}
	}

	/*
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
							g.mainChar.posX < float64(x)+1 {
							if g.mainChar.posY >= float64(y)-1 &&
								g.mainChar.posY < float64(y)+0.5 {
								op := &ebiten.DrawImageOptions{}
								op.GeoM.Scale(3, 3)
								op.GeoM.Translate(g.mainChar.posX*256+g.mainChar.posY*256+xshift+50, g.mainChar.posY*154-g.mainChar.posX*144-float64(g.mainChar.posZ*120)+yshift-370)

								screen.DrawImage(g.mainChar.currentImage, op)
							}
						}
					}
				}
			}
		}
	*/

	return
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 6400, 3200
}
