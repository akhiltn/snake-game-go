package game

import (
	"bytes"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

const GameSpeed = time.Second / 6

var gameFont ebitentext.Face

func init() {
	src, err := ebitentext.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		panic(err)
	}
	gameFont = &ebitentext.GoTextFace{Source: src, Size: 24}
}

type GameImage struct {
	SnakeImg *ebiten.Image
	FoodImg  *ebiten.Image
}

type Game struct {
	snake      *Snake
	direction  Direction
	lastUpdate time.Time
	food       Food
	started    bool
	gameOver   bool
	GameImage
}

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

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *NewGame()
		}
		return nil
	}

	if !g.started {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.started = true
			g.lastUpdate = time.Now()
		}
		return nil
	}

	if time.Since(g.lastUpdate) < GameSpeed {
		return nil
	}

	g.HandleInput()

	head := g.snake.NextHead(g.direction)
	g.gameOver = g.snake.WillEatSelf(head)
	if !g.gameOver {
		g.snake.MoveHead(g.direction)
		if g.snake.Head() != Point(g.food) {
			g.snake.MoveTail()
		} else {
			g.food = SpawnFood()
		}
	}

	g.lastUpdate = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		g.drawGameOver(screen)
		return
	}

	if !g.started {
		g.drawStartScreen(screen)
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(g.food.X*PixelSize),
		float64(g.food.Y*PixelSize),
	)
	screen.DrawImage(g.FoodImg, op)

	for _, p := range g.snake.Body() {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.X*PixelSize), float64(p.Y*PixelSize))
		screen.DrawImage(g.SnakeImg, op)
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	op := &ebitentext.DrawOptions{}
	op.GeoM.Translate(float64(ScreenWidth)/2, float64(ScreenHeight)/2)
	op.PrimaryAlign = ebitentext.AlignCenter
	op.SecondaryAlign = ebitentext.AlignCenter
	op.LineSpacing = 16
	ebitentext.Draw(screen, "GAME OVER\n\nPress R to restart", gameFont, op)
}

func (g *Game) drawStartScreen(screen *ebiten.Image) {
	op := &ebitentext.DrawOptions{}
	op.GeoM.Translate(float64(ScreenWidth)/2, float64(ScreenHeight)/2)
	op.PrimaryAlign = ebitentext.AlignCenter
	op.SecondaryAlign = ebitentext.AlignCenter
	op.LineSpacing = 16
	ebitentext.Draw(screen, "SNAKE GAME\n\nPress Enter to start", gameFont, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	centerX := (ScreenWidth / PixelSize) / 2
	centerY := (ScreenHeight / PixelSize) / 2

	snakeImg := ebiten.NewImage(PixelSize, PixelSize)
	snakeImg.Fill(color.RGBA{0, 255, 0, 255})

	foodImg := ebiten.NewImage(PixelSize, PixelSize)
	foodImg.Fill(color.RGBA{255, 0, 0, 255})

	return &Game{
		snake:      NewSnake(Point{X: centerX, Y: centerY}),
		direction:  Right,
		lastUpdate: time.Now(),
		food:       SpawnFood(),
		started:    false,
		gameOver:   false,
		GameImage: GameImage{
			SnakeImg: snakeImg,
			FoodImg:  foodImg,
		},
	}
}
