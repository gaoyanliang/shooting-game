package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Bullet struct {
	GameObject
	image *ebiten.Image

	// 子弹移动系数
	speedFactor float64
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		image: img,
		GameObject: GameObject{
			width:  cfg.BulletWidth,
			height: cfg.BulletHeight,
			// 子弹从飞船头部中间射出
			x: ship.x + float64(ship.width-cfg.BulletWidth)/2,
			y: float64(cfg.ScreenHeight - ship.height - cfg.BulletHeight),
		},
		speedFactor: cfg.BulletSpeedFactor,
	}
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}

func (bullet *Bullet) outOfScreen() bool {
	return bullet.y < -float64(bullet.height)
}
