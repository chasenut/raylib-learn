package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type timerInterface interface {
	float32
	float32
	GetStatus() bool
	Init(float32)
}

type timer struct {
	lifeTime float32
	initTime float32
}

type repeater struct {
	lifeTime float32
	initTime float32
}

func (t *timer) Init(lifeTime float32) {
	t.initTime = float32(rl.GetTime())
	t.lifeTime = lifeTime
}	

func (r *repeater) Init(lifeTime float32) {
	r.initTime = float32(rl.GetTime())
	r.lifeTime = lifeTime
}	

func (t *timer) GetStatus() bool {
	if float32(rl.GetTime()) - t.initTime > t.lifeTime {
		return true
	}
	return false
}

func (r *repeater) GetStatus() bool {
	if float32(rl.GetTime()) - r.initTime > r.lifeTime {
		r.initTime = float32(rl.GetTime())
		return true
	}
	return false
}

const (
	screenWidth int32 = 1000
	screenHeight int32 = 480

	FPS int32 = 60
)

var (
	dt float32 = rl.GetFrameTime()
	animTimer repeater = repeater{}
	animFrame bool

	running bool = true
	bkgColor rl.Color = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle
	playerMoving bool
	playerDir int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame int

	playerSpeed float32 = 150

	frameCount int

	musicPaused bool
	music rl.Music

	cam rl.Camera2D
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerUp = true
		playerDir = 1 // for spritesheet anims
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDown = true
		playerDir = 0 // for spritesheet anims
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerRight = true
		playerDir = 3 // for spritesheet anims
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerLeft = true
		playerDir = 2 // for spritesheet anims
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func update() {
	running = !rl.WindowShouldClose()
	dt = rl.GetFrameTime()
	animFrame = animTimer.GetStatus()

	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed * dt
		}
		if playerDown {
			playerDest.Y += playerSpeed * dt
		}
		if playerRight {
			playerDest.X += playerSpeed * dt
		}
		if playerLeft {
			playerDest.X -= playerSpeed * dt
		}

	} 
	if animFrame {
		playerFrame++
	}
	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}	
	if !playerMoving && playerFrame > 1 {
		playerFrame = 0
	}

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	playerSrc.Y = playerSrc.Height * float32(playerDir)
	
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(
			float32(playerDest.X-(playerDest.Width/2)), 
			float32(playerDest.Y-(playerDest.Height/2)),
		)

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib-learn-go")
	rl.SetExitKey(0)
	rl.SetTargetFPS(FPS)

	animTimer.Init(.5)

	grassSprite = rl.LoadTexture("res/tilesets/grass.png")
	playerSprite = rl.LoadTexture("res/characters/character-spritesheet.png")
	
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/audio/bkg.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(
		rl.NewVector2(
			float32(screenWidth/2), 
			float32(screenHeight/2)), 
		rl.NewVector2(
			float32(playerDest.X-(playerDest.Width/2)), 
			float32(playerDest.Y-(playerDest.Height/2))), 
		0.0, 
		1.5,
		)
}

func quit() {
	defer rl.CloseWindow()

	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
}

func main() {
	for running {
		input()
		update()
		render()
	}
	quit()
}
