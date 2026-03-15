package game

import "math/rand/v2"

type Food Point

func SpawnFood(snake *Snake) Food {
	body := make(map[Point]bool)
	for _, p := range snake.Body() {
		body[p] = true
	}

	for {
		food := Food{
			X: rand.IntN(ScreenWidth / PixelSize),
			Y: rand.IntN(ScreenHeight / PixelSize),
		}
		if !body[Point(food)] {
			return food
		}
	}
}
