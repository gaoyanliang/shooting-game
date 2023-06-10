package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Monster struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewMonster(cfg *Config) *Monster {
	img, _, err := ebitenutil.NewImageFromFile("./images/monster.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	return &Monster{
		image: img,
		GameObject: GameObject{
			width:  width,
			height: height,
			x:      0,
			y:      0,
		},
		speedFactor: cfg.MonsterSpeedFactor,
	}
}

func (monster *Monster) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(monster.x, monster.y)
	screen.DrawImage(monster.image, op)
}

func (monster *Monster) outOfScreen(cfg *Config) bool {
	return monster.y > float64(cfg.ScreenHeight)
}
