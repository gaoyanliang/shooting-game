package main

//
//import (
//	"log"
//
//	"github.com/hajimehoshi/ebiten/v2"
//	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
//)
//
//// 仅用来展示窗口，暂不定义游戏数据
//type Game struct{}
//
//func (g *Game) Update() error {
//	return nil
//}
//
//func (g *Game) Draw(screen *ebiten.Image) {
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
