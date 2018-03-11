package liskov

import "testing"

func TestSquare(t *testing.T) {
	s := &Square{}
	testBase(t, s)
}
