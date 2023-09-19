package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	heroSprite rl.Texture2D

	spriteX      float32 = (screenWidth / 2) + (spriteWidth / 2)
	spriteY      float32 = (screenHeight / 2) + (spriteHeight / 2)
	spriteWidth  float32 = 80
	spriteHeight float32 = 80

	cam rl.Camera2D
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Second game")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	heroSprite = rl.LoadTexture("./res/hollownight1.png")

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32((spriteX-spriteWidth/2)), float32(spriteY-(spriteHeight/2))), 0.0, 1.5)

}

func quit() {
	rl.UnloadTexture(heroSprite)
	rl.CloseWindow()
}

func main() {
	p := NewPlayer(heroSprite, rl.NewRectangle(0, 2, 82, 80), rl.NewRectangle(spriteX, spriteY, spriteWidth, spriteHeight), 2)

	g := NewGame(rl.NewColor(90, 90, 90, 1), cam)

	for !rl.WindowShouldClose() {
		p.input()
		g.update(&p)
		g.render(&p)
	}
}
