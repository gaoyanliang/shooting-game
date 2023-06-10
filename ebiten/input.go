package main

//
//import (
//	"fmt"
//	"github.com/hajimehoshi/ebiten/v2"
//	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
//	"github.com/hajimehoshi/ebiten/v2/inpututil"
//	"image/color"
//	"log"
//)
//
//// 仅用来展示窗口，暂不定义游戏数据
//type Game struct{}
//
//func (g *Game) Update() error {
//	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
//		fmt.Println("左键 ⬅️")
//	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
//		fmt.Println("右键 ➡️️")
//	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
//		fmt.Println("下键 ⬇️️")
//	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
//		fmt.Println("上键 ⬆️️")
//	} else {
//		// ... 暂不处理
//	}
//	return nil
//}
//
//func (g *Game) Draw(screen *ebiten.Image) {
//	screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})
//	ebitenutil.DebugPrint(screen, "Hello, World")
//}
//
//func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return 100, 100
//}
//
//func main() {
//	ebiten.SetWindowSize(200, 200)
//	ebiten.SetWindowTitle("生成游戏窗口")
//	if err := ebiten.RunGame(&Game{}); err != nil {
//		log.Fatal(err)
//	}
//}
