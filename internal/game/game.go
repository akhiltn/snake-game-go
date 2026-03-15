package game

import (
	"errors"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameImage struct {
	SnakeImg *ebiten.Image
	HeadImg  *ebiten.Image
	FoodImg  *ebiten.Image
}

type Game struct {
	snake      *Snake
	direction  Direction
	lastUpdate time.Time
	food       Food
	started    bool
	gameOver   bool
	paused     bool
	quit       bool
	GameImage
}

func (g *Game) Update() error {
	g.handleStateInput()

	if g.quit {
		return errors.New("quit")
	}

	if g.gameOver || !g.started || g.paused {
		return nil
	}

	g.HandleInput()

	if time.Since(g.lastUpdate) < GameSpeed {
		return nil
	}

	head := g.snake.NextHead(g.direction)
	g.gameOver = g.snake.WillEatSelf(head)
	if !g.gameOver {
		g.snake.MoveHead(g.direction)
		if g.snake.Head() != Point(g.food) {
			g.snake.MoveTail()
		} else {
			g.food = SpawnFood(g.snake)
		}
	}

	g.lastUpdate = time.Now()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
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

	if g.paused {
		g.drawPaused(screen)
		return
	}

	g.drawGame(screen)
}

func NewGame() *Game {
	centerX := (ScreenWidth / PixelSize) / 2
	centerY := (ScreenHeight / PixelSize) / 2

	snake := NewSnake(Point{X: centerX, Y: centerY})

	snakeImg := ebiten.NewImage(PixelSize, PixelSize)
	snakeImg.Fill(snakeColor)

	headImg := ebiten.NewImage(PixelSize, PixelSize)
	headImg.Fill(headColor)

	foodImg := ebiten.NewImage(PixelSize, PixelSize)
	foodImg.Fill(foodColor)

	return &Game{
		snake:      snake,
		direction:  Right,
		lastUpdate: time.Now(),
		food:       SpawnFood(snake),
		started:    false,
		paused:     false,
		gameOver:   false,
		GameImage: GameImage{
			SnakeImg: snakeImg,
			HeadImg:  headImg,
			FoodImg:  foodImg,
		},
	}
}
