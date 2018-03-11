package liskov

// Rectangle ...
type Rectangle struct {
	width  int
	height int
}

// SetWidth ...
func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}

// SetHeight ...
func (r *Rectangle) SetHeight(newHeight int) {
	r.height = newHeight
}

// GetArea ...
func (r *Rectangle) GetArea() int {
	return r.width * r.height
}
