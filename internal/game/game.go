package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const GameSpeed = time.Second / 6

type GameImage struct {
	SnakeImg    *ebiten.Image
	FoodImg     *ebiten.Image
	GameOverImg *ebiten.Image
}

type Game struct {
	snake      *Snake
	direction  Direction
	lastUpdate time.Time
	food       Food
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
	if time.Since(g.lastUpdate) < GameSpeed {
		return nil
	}

	if g.gameOver {
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
		gameOver:   false,
		GameImage: GameImage{
			SnakeImg: snakeImg,
			FoodImg:  foodImg,
		},
	}
}
