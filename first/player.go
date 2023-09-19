package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Sprite       rl.Texture2D
	Source       rl.Rectangle
	Destination  rl.Rectangle
	Speed        float32
	Moving       bool
	Left, Right  bool
	Jump, Crouch bool
	Frames       int
}

func NewPlayer(sprite rl.Texture2D, src rl.Rectangle, dest rl.Rectangle, speed float32) Player {
	return Player{
		Sprite:      sprite,
		Source:      src,
		Destination: dest,
		Speed:       speed,
		Moving:      false,
		Left:        false,
		Right:       false,
		Jump:        false,
		Crouch:      false,
		Frames:      0,
	}
}

func (p *Player) resetMotion() {
	p.Moving = false
	p.Left = false
	p.Right = false
	p.Jump = false
	p.Crouch = false
}

func (p *Player) update(g *Game) {
	if p.Moving {
		if p.Right {
			if p.Source.Width < 0 {
				p.Source.Width = -p.Source.Width
			}
		}
		if p.Left {
			if p.Source.Width > 0 {
				p.Source.Width = -p.Source.Width
			}
		}
		if g.Frames%8 == 1 {
			p.Frames++
		}
	} else if g.Frames%45 == 1 {
		p.Frames++
	}

	if p.Frames > 5 {
		p.Frames = 0
	}

	if !p.Moving && p.Frames > 1 {
		p.Frames = 0
	}

	if p.Source.Width < 0 {
		p.Source.X = -p.Source.Width * float32(p.Frames)
	} else {
		p.Source.X = p.Source.Width * float32(p.Frames)
	}

	p.resetMotion()
}

func (p *Player) input() {
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		p.Moving = true
		p.Right = true
		p.Destination.X += p.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		p.Moving = true
		p.Left = true
		p.Destination.X -= p.Speed
	}
}