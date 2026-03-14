package game

import "math/rand/v2"

type Food Point

func SpawnFood() Food {
	return Food{
		X: rand.IntN(ScreenWidth / PixelSize),
		Y: rand.IntN(ScreenHeight / PixelSize),
	}
}
