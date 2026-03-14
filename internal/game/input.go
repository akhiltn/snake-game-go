package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.direction != Down {
		g.direction = Up
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.direction != Up {
		g.direction = Down
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.direction != Right {
		g.direction = Left
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.direction != Left {
		g.direction = Right
	}
}

func (g *Game) handleStateInput() {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *NewGame()
		}
		return
	}

	if !g.started {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.started = true
			g.lastUpdate = time.Now()
		}
		return
	}

	if g.paused {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.paused = false
		}
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = true
	}
}
