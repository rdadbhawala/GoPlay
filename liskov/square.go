package liskov

// Square ...
type Square struct {
	Rectangle
}

// SetWidth ...
func (s *Square) SetWidth(newWidth int) {
	s.setSide(newWidth)
}

// SetHeight ...
func (s *Square) SetHeight(newHeight int) {
	s.setSide(newHeight)
}

func (s *Square) setSide(newSide int) {
	s.Rectangle.SetWidth(newSide)
	s.Rectangle.SetHeight(newSide)
}
