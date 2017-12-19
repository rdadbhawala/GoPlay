package interfaceIsPointer

import "testing"

type number interface {
	GetNumber() int
	SetNumber(newNum int)
}

type actualNumber struct {
	num int
}

func (a actualNumber) GetNumber() int {
	return a.num
}

// if the following line changes the pointer receiver to struct receiver, the test fails
func (a *actualNumber) SetNumber(newNum int) {
	a.num = newNum
}

func TestInterfaceIsPointer(t *testing.T) {
	a := actualNumber{
		num: 100,
	}
	a.SetNumber(50)
	var i number
	i = &a
	check(i, 200)
	if a.num != 200 {
		t.Error("Interface IS NOT a pointer", a.num)
	}
}

func check(inum number, newValue int) {
	inum.SetNumber(newValue)
}
