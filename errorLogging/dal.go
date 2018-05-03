package errorLogging

import (
	"errors"
	"time"
)

func newDal(pErr1 bool, pErr2 bool) dal {
	return &dalImpl{pErr1, pErr2}
}

type dalImpl struct {
	err1 bool
	err2 bool
}

func (d *dalImpl) Op1() error {
	time.Sleep(time.Millisecond * 10)
	if d.err1 {
		return errors.New("err1")
	}
	return nil
}

func (d *dalImpl) Op1Panic() {
	time.Sleep(time.Millisecond * 10)
	if d.err1 {
		panic(errors.New("err1"))
	}
}

func (d *dalImpl) Op2() error {
	time.Sleep(time.Millisecond * 10)
	if d.err2 {
		return errors.New("err2")
	}
	return nil
}

func (d *dalImpl) Op2Panic() {
	time.Sleep(time.Millisecond * 10)
	if d.err2 {
		panic(errors.New("err2"))
	}
}
