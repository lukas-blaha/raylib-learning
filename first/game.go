package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Background rl.Color
	Camera     rl.Camera2D
	Ground     float32
	Gravity    float32
	Frames     int
}

func NewGame(bkg rl.Color, cam rl.Camera2D, ground, gravity float32) Game {
	return Game{
		Background: bkg,
		Camera:     cam,
		Ground:     ground,
		Gravity:    gravity,
		Frames:     0,
	}
}

func (g *Game) drawScene(p *Player) {
	rl.DrawTexturePro(p.Sprite, p.Source, p.Destination, rl.NewVector2(p.Destination.Width, p.Destination.Height), 0, rl.White)
}

func (g *Game) update(p *Player) {
	g.Frames++
}

func (g *Game) render(p *Player) {
	rl.BeginDrawing()
	rl.ClearBackground(g.Background)
	rl.BeginMode2D(g.Camera)

	p.update(g)
	g.drawScene(p)

	rl.EndMode2D()
	rl.EndDrawing()
}
