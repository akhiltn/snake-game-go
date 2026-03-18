package game

type Point struct {
	X, Y int
}

// WrapToGrid bounds the point's coordinates within a wrapping (toroidal) grid.
func (p Point) WrapToGrid(maxX, maxY int) Point {
	return Point{
		X: (p.X + maxX) % maxX,
		Y: (p.Y + maxY) % maxY,
	}
}
