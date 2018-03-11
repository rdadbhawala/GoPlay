package alternate

import "github.com/rdadbhawala/GoPlay/liskov"

// NewSquare ...
type NewSquare struct {
	r liskov.Rectangle
}

// SetSide ...
func (n *NewSquare) SetSide(newSide int) {
	n.r.SetWidth(newSide)
	n.r.SetHeight(newSide)
}

// GetArea ....
func (n *NewSquare) GetArea() int {
	return n.r.GetArea()
}
