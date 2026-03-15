package game

import "math/rand/v2"

type Food Point

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
