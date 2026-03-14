package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (g *Game) drawGame(screen *ebiten.Image) {
	g.drawFood(screen)
	g.drawSnake(screen)
}

func (g *Game) drawFood(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(g.food.X*PixelSize),
		float64(g.food.Y*PixelSize),
	)
	screen.DrawImage(g.FoodImg, op)
}

func (g *Game) drawSnake(screen *ebiten.Image) {
	for _, p := range g.snake.Body() {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.X*PixelSize), float64(p.Y*PixelSize))
		screen.DrawImage(g.SnakeImg, op)
	}
}

func (g *Game) drawCenteredText(screen *ebiten.Image, text string) {
	op := &ebitentext.DrawOptions{}
	op.GeoM.Translate(float64(ScreenWidth)/2, float64(ScreenHeight)/2)
	op.PrimaryAlign = ebitentext.AlignCenter
	op.SecondaryAlign = ebitentext.AlignCenter
	op.LineSpacing = 16
	ebitentext.Draw(screen, text, gameFont, op)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	g.drawCenteredText(screen, "GAME OVER\n\nPress R to restart")
}

func (g *Game) drawStartScreen(screen *ebiten.Image) {
	g.drawCenteredText(screen, "SNAKE GAME\n\nPress Enter to start")
}

func (g *Game) drawPaused(screen *ebiten.Image) {
	g.drawCenteredText(screen, "PAUSED\n\nPress Space to resume")
}
