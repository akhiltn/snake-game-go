package game

type Snake struct {
	body  []Point
	cache map[Point]bool
}

func NewSnake(start Point) *Snake {
	return &Snake{
		body:  []Point{start},
		cache: map[Point]bool{start: true},
	}
}

func (s *Snake) Head() Point {
	return s.body[len(s.body)-1]
}

func (s *Snake) NextHead(d Direction) Point {
	head := s.Head()

	switch d {
	case Up:
		head.Y--
	case Down:
		head.Y++
	case Left:
		head.X--
	case Right:
		head.X++
	}

	head.X = wrap(head.X, ScreenWidth/PixelSize)
	head.Y = wrap(head.Y, ScreenHeight/PixelSize)

	return head
}

func (s *Snake) Body() []Point {
	return s.body
}

func (s *Snake) MoveHead(d Direction) {
	head := s.NextHead(d)
	s.body = append(s.body, head)
	s.cache[head] = true
}

func (s *Snake) MoveTail() {
	tail := s.body[0]
	delete(s.cache, tail)
	s.body = s.body[1:]
}

func wrap(v, max int) int {
	return (v + max) % max
}

func (s *Snake) WillEatSelf(next Point) bool {
	return s.cache[next]
}
