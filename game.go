package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"time"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeOver
)

type Game struct {
	mode     Mode
	ship     *Ship                 // 飞船
	bullets  map[*Bullet]struct{}  // 子弹
	monsters map[*Monster]struct{} // 怪物 👹
	cfg      *Config
}

func (g *Game) init() {
	g.CreateMonsters()
	g.CreateFonts()
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	g := &Game{
		ship:     NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		bullets:  make(map[*Bullet]struct{}),
		monsters: make(map[*Monster]struct{}),
		cfg:      cfg,
	}
	g.init()
	return g
}

func (g *Game) CreateMonsters() {
	monster := NewMonster(g.cfg)

	availableSpaceX := g.cfg.ScreenWidth - 2*monster.width
	numMonsters := availableSpaceX / (2 * monster.width)

	for row := 0; row < 2; row++ {
		for i := 0; i < numMonsters; i++ {
			monster = NewMonster(g.cfg)
			monster.x = float64(monster.width + 2*monster.width*i)
			monster.y = float64(monster.height*row) * 1.5
			g.addMonster(monster)
		}
	}
}

func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

func (g *Game) addMonster(monster *Monster) {
	g.monsters[monster] = struct{}{}
}

// CheckCollision 检查子弹和外星人之间是否有碰撞
func CheckCollision(bullet *Bullet, monster *Monster) bool {
	monsterTop, monsterLeft := monster.y, monster.x
	monsterBottom, monsterRight := monster.y+float64(monster.height), monster.x+float64(monster.width)
	// 左上角
	x, y := bullet.x, bullet.y
	if y > monsterTop && y < monsterBottom && x > monsterLeft && x < monsterRight {
		return true
	}

	// 右上角
	x, y = bullet.x+float64(bullet.width), bullet.y
	if y > monsterTop && y < monsterBottom && x > monsterLeft && x < monsterRight {
		return true
	}

	// 左下角
	x, y = bullet.x, bullet.y+float64(bullet.height)
	if y > monsterTop && y < monsterBottom && x > monsterLeft && x < monsterRight {
		return true
	}

	// 右下角
	x, y = bullet.x+float64(bullet.width), bullet.y+float64(bullet.height)
	if y > monsterTop && y < monsterBottom && x > monsterLeft && x < monsterRight {
		return true
	}

	return false
}

func (g *Game) CheckCollision() {
	for monster := range g.monsters {
		for bullet := range g.bullets {
			if CheckCollision(bullet, monster) {
				delete(g.monsters, monster)
				delete(g.bullets, bullet)
			}
		}
	}
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		// 左键 或 鼠标左键 开始游戏
		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.mode = ModeGame
		}
	case ModeGame:
		for bullet := range g.bullets {
			bullet.y -= bullet.speedFactor
		}

		for monster := range g.monsters {
			monster.y += monster.speedFactor
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			g.ship.x -= g.cfg.ShipSpeedFactor
			if g.ship.x < -float64(g.ship.width)/2 {
				g.ship.x = -float64(g.ship.width) / 2
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			g.ship.x += g.cfg.ShipSpeedFactor
			if g.ship.x > float64(g.cfg.ScreenWidth)-float64(g.ship.width)/2 {
				g.ship.x = float64(g.cfg.ScreenWidth) - float64(g.ship.width)/2
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			if len(g.bullets) < g.cfg.MaxBulletNum &&
				time.Now().Sub(g.ship.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
				bullet := NewBullet(g.cfg, g.ship)
				g.addBullet(bullet)
				g.ship.lastBulletTime = time.Now()
			}
		}

		g.CheckCollision()

		for bullet := range g.bullets {
			if bullet.outOfScreen() {
				delete(g.bullets, bullet)
			}
		}
	case ModeOver:
		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.init()
			g.mode = ModeTitle
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)

	// 绘制界面字体
	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"SHOOTING GAME"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		g.ship.Draw(screen)
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		for monter := range g.monsters {
			monter.Draw(screen)
		}
	case ModeOver:
		texts = []string{"", "GAME OVER!"}
	}

	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.cfg.TitleFontSize, color.White)
	}
	for i, l := range texts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.cfg.FontSize, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}
