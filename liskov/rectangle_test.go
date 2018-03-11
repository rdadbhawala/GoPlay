package liskov

import "testing"

func TestRectangle(t *testing.T) {
	r := &Rectangle{}
	testBase(t, r)
}
