package game

import (
	"errors"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	snake      *Snake
	direction  Direction
	lastUpdate time.Time
	food       Food
	started    bool
	gameOver   bool
	paused     bool
	quit       bool
	renderer   *Renderer
}

var _ ebiten.Game = (*Game)(nil)

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
			g.food = g.SpawnFood()
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
		g.renderer.DrawGameOver(screen)
		return
	}

	if !g.started {
		g.renderer.DrawStartScreen(screen)
		return
	}

	if g.paused {
		g.renderer.DrawPaused(screen)
		return
	}

	g.renderer.DrawGrid(screen)
	g.renderer.DrawFood(screen, g.food)
	g.renderer.DrawSnake(screen, g.snake)
}

func (g *Game) SpawnFood() Food {
	for {
		food := Food{
			X: rand.IntN(ScreenWidth / PixelSize),
			Y: rand.IntN(ScreenHeight / PixelSize),
		}
		if !g.snake.Contains(Point(food)) {
			return food
		}
	}
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

	renderer := NewRenderer(snakeImg, headImg, foodImg)

	g := &Game{
		snake:      snake,
		direction:  Right,
		lastUpdate: time.Now(),
		started:    false,
		paused:     false,
		gameOver:   false,
		renderer:   renderer,
	}
	g.food = g.SpawnFood()
	return g
}

