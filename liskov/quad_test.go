package liskov

import "testing"

func testBase(t *testing.T, q Quad) {
	q.SetWidth(5)
	q.SetHeight(10)
	if q.GetArea() != 50 {
		t.Error("Area was not 50: ", q.GetArea())
	}
}
