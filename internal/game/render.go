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

type Renderer struct {
	SnakeImg *ebiten.Image
	HeadImg  *ebiten.Image
	FoodImg  *ebiten.Image
}

func NewRenderer(snakeImg, headImg, foodImg *ebiten.Image) *Renderer {
	return &Renderer{
		SnakeImg: snakeImg,
		HeadImg:  headImg,
		FoodImg:  foodImg,
	}
}

func (r *Renderer) DrawGrid(screen *ebiten.Image) {
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

func (r *Renderer) DrawFood(screen *ebiten.Image, food Food) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(food.X*PixelSize),
		float64(food.Y*PixelSize),
	)
	screen.DrawImage(r.FoodImg, op)
}

func (r *Renderer) DrawSnake(screen *ebiten.Image, snake *Snake) {
	body := snake.Body()
	for i := len(body) - 1; i >= 0; i-- {
		p := body[i]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.X*PixelSize), float64(p.Y*PixelSize))
		if i == len(body)-1 {
			screen.DrawImage(r.HeadImg, op)
		} else {
			screen.DrawImage(r.SnakeImg, op)
		}
	}
}

func (r *Renderer) drawCenteredText(screen *ebiten.Image, text string) {
	op := &ebitentext.DrawOptions{}
	op.GeoM.Translate(float64(ScreenWidth)/2, float64(ScreenHeight)/2)
	op.PrimaryAlign = ebitentext.AlignCenter
	op.SecondaryAlign = ebitentext.AlignCenter
	op.LineSpacing = 16
	ebitentext.Draw(screen, text, gameFont, op)
}

func (r *Renderer) DrawGameOver(screen *ebiten.Image) {
	r.drawCenteredText(screen, "GAME OVER\n\nPress R to restart\n\nPress Q to quit")
}

func (r *Renderer) DrawStartScreen(screen *ebiten.Image) {
	r.drawCenteredText(screen, "SNAKE GAME\n\nPress Enter to start\n\nPress Q to quit")
}

func (r *Renderer) DrawPaused(screen *ebiten.Image) {
	r.drawCenteredText(screen, "PAUSED\n\nPress Space to resume\n\nPress Q to quit")
}
