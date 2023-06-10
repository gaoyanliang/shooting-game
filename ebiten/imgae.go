package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

const (
	screenWidth  = 600
	screenHeight = 500
	title        = "射击游戏"
)

// 仅用来展示窗口，暂不定义游戏数据
type Game struct {
	ship *Ship
}

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x      float64 // x坐标
	y      float64 // y坐标
}

func NewShip() *Ship {
	img, _, err := ebitenutil.NewImageFromFile("/Users/gaoyanliang/github/play-with-pulsar/game/images/ship.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight-height) / 2,
	}

	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
}

func NewGame() *Game {
	return &Game{
		ship: NewShip(),
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		fmt.Println("左键 ⬅️")
		g.ship.x -= 10 // 移动范围太小不容易出效果
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		fmt.Println("右键 ➡️️")
		g.ship.x += 10
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		fmt.Println("下键 ⬇️️")
		g.ship.y -= 10
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		fmt.Println("上键 ⬆️️")
		g.ship.y += 10
	} else {
		// ... 暂不处理
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})
	g.ship.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
