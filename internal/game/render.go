package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	lightColor = color.RGBA{235, 235, 235, 255}
	darkColor  = color.RGBA{215, 215, 215, 255}
)

var (
	snakeColor = color.RGBA{46, 204, 113, 255}
	headColor  = color.RGBA{39, 174, 96, 255}
	foodColor  = color.RGBA{231, 76, 60, 255}
)

func (g *Game) drawGame(screen *ebiten.Image) {
	g.drawGrid(screen)
	g.drawFood(screen)
	g.drawSnake(screen)
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	for gx := 0; gx < ScreenWidth/PixelSize; gx++ {
		for gy := 0; gy < ScreenHeight/PixelSize; gy++ {
			var c color.RGBA
			if (gx+gy)%2 == 0 {
				c = lightColor
			} else {
				c = darkColor
			}
			for x := gx * PixelSize; x < (gx+1)*PixelSize; x++ {
				for y := gy * PixelSize; y < (gy+1)*PixelSize; y++ {
					screen.Set(x, y, c)
				}
			}
		}
	}
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
	body := g.snake.Body()
	for i := len(body) - 1; i >= 0; i-- {
		p := body[i]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.X*PixelSize), float64(p.Y*PixelSize))
		if i == len(body)-1 {
			screen.DrawImage(g.HeadImg, op)
		} else {
			screen.DrawImage(g.SnakeImg, op)
		}
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
	g.drawCenteredText(screen, "GAME OVER\n\nPress R to restart\n\nPress Q to quit")
}

func (g *Game) drawStartScreen(screen *ebiten.Image) {
	g.drawCenteredText(screen, "SNAKE GAME\n\nPress Enter to start\n\nPress Q to quit")
}

func (g *Game) drawPaused(screen *ebiten.Image) {
	g.drawCenteredText(screen, "PAUSED\n\nPress Space to resume\n\nPress Q to quit")
}
