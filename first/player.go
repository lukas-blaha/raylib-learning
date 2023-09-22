package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Sprite       rl.Texture2D
	Moves        map[string][3]int
	Source       rl.Rectangle
	Destination  rl.Rectangle
	Speed        float32
	JumpHeight   int
	Moving       bool
	Left, Right  bool
	Jump, Crouch bool
	Attack       bool
	Frames       int
}

func NewPlayer(sprite rl.Texture2D, moves map[string][3]int, src rl.Rectangle, dest rl.Rectangle, speed float32) Player {
	return Player{
		Sprite:      sprite,
		Source:      src,
		Moves:       moves,
		Destination: dest,
		Speed:       speed,
		JumpHeight:  5,
		Moving:      false,
		Left:        false,
		Right:       false,
		Jump:        false,
		Crouch:      false,
		Attack:      false,
		Frames:      0,
	}
}

func (p *Player) resetMotion() {
	p.Moving = false
	p.Left = false
	p.Right = false
	p.Jump = false
	p.Crouch = false
	p.Attack = false
}

func (p *Player) getFrames(move string) (int, int, int) {
	return p.Moves[move][0], p.Moves[move][1], p.Moves[move][2]
}

func (p *Player) update(g *Game) {
	var y, sx, ex int
	baseY := float32(0 + p.Sprite.Height)

	if p.Moving {
		if p.Right {
			if p.Source.Width < 0 {
				p.Source.Width = -p.Source.Width
			}
			y, sx, ex = p.getFrames("go")
		}
		if p.Left {
			if p.Source.Width > 0 {
				p.Source.Width = -p.Source.Width
			}
			y, sx, ex = p.getFrames("go")
		}
		if p.Jump {
			y, sx, ex = p.getFrames("jump")
			if p.JumpHeight > 0 {
				p.Destination.Y -= p.Speed
			} else if p.Destination.Y > baseY {
				p.Destination.Y += float32(g.Gravity)
			} else {
				p.JumpHeight = 6
			}
			p.JumpHeight--
		}
		if p.Attack {
			y, sx, ex = p.getFrames("attack")
		}
		if g.Frames%8 == 1 {
			p.Frames++
		}
	} else if g.Frames%45 == 1 {
		p.Frames++
	}

	if !p.Moving {
		y, sx, ex = p.getFrames("idle")
	}

	p.Source.Y = (p.Source.Height * float32(y)) + 2

	if p.Frames < sx {
		p.Frames = sx
	}
	if p.Frames == ex {
		p.Frames = sx
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
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		p.Moving = true
		p.Jump = true
	}
	if rl.IsKeyDown(rl.KeySpace) {
		p.Moving = true
		p.Attack = true
	}
}
