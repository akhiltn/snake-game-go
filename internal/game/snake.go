package game

type Snake []Point

func (s *Snake) MoveTail() {
	if len(*s) > 1 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *Snake) MoveHead(d Direction) {
	head := (*s)[0]

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

	*s = append([]Point{head}, (*s)...)
}
