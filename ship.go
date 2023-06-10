package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

type Ship struct {
	image          *ebiten.Image
	width          int     // 飞船宽度
	height         int     // 飞船高度
	x              float64 // x 坐标
	y              float64 // y 坐标
	lastBulletTime time.Time
}

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./images/rocket.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}

	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
}
