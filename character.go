package main

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
)

var characterImage *ebiten.Image

const (
	idle int = iota
	walking
	numMoveTypes
)

type character struct {
	posX         float64
	posY         float64
	posZ         int
	direction    int
	moveType     int
	frame        int
	animStep     int
	currentImage *ebiten.Image
	allImages    [][][]*ebiten.Image
}

func (c *character) updateAnim() {
	c.frame = (c.frame + 1) % 60
	c.animStep = c.frame / 10
	c.currentImage = c.allImages[c.moveType][c.direction][c.animStep]
}

func makeEtna(x, y float64, z int) character {
	c := character{
		posX: x,
		posY: y,
		posZ: z,
	}
	c.direction = 0
	c.moveType = idle
	c.frame = 0
	c.animStep = 0
	c.allImages = make([][][]*ebiten.Image, numMoveTypes)
	for move := 0; move < numMoveTypes; move++ {
		c.allImages[move] = make([][]*ebiten.Image, 4)
		for direction := 0; direction < 4; direction++ {
			c.allImages[move][direction] = make([]*ebiten.Image, 6)
			for step := 0; step < 6; step++ {
				sx := step * 128
				sy := (move*4 + direction) * 192
				c.allImages[move][direction][step] = characterImage.SubImage(
					image.Rect(
						sx, sy,
						sx+128, sy+192,
					),
				).(*ebiten.Image)
			}
		}
	}
	return c
}

func (g *game) moveChar() bool {
	var dX, dY float64
	switch g.mainChar.direction {
	case 0:
		dX = -0.015
	case 1:
		dY = -0.015
	case 2:
		dX = 0.015
	case 3:
		dY = 0.015
	}
	newX := g.mainChar.posX + dX
	newY := g.mainChar.posY + dY
	if newX < 0 || newY < 0 ||
		newX >= float64(len(g.field[0][0]))-1 ||
		newY >= float64(len(g.field[0]))-1 {
		return false
	}
	newXInt := int(math.Round(newX))
	newYInt := int(math.Round(newY))
	fmt.Println(newXInt, newYInt)
	z := g.mainChar.posZ
	if z+2 < len(g.field) && g.field[z+2][newYInt][newXInt] != nilTile {
		return false
	}
	if z+1 < len(g.field) && g.field[z+1][newYInt][newXInt] != nilTile {
		g.mainChar.posZ++
	} else if g.field[z][newYInt][newXInt] == nilTile {
		if z-1 >= 0 && g.field[z-1][newYInt][newXInt] == nilTile {
			return false
		}
		g.mainChar.posZ--
	}
	g.mainChar.posX = newX
	g.mainChar.posY = newY
	return true
}
