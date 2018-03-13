package errorReturn

import (
	"errors"
	"testing"
)

const (
	anError = "anError"
)

func returnError() error {
	return nil // errors.New(anError)
}

func returnError2() error {
	return errors.New(anError)
}

func returnFunc() error {
	err := returnError()
	if err != nil {
		return err
	}

	err = returnError()
	if err != nil {
		return err
	}

	err = returnError2()
	if err != nil {
		return err
	}

	return nil
}

func BenchmarkReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		returnFunc()
	}
}

func panicError() {
	return // nil // errors.New(anError)
}

func panicError2() {
	panic(errors.New(anError))
}

func panicFunc() {
	defer func() {
		recover()
	}()
	panicError()
	panicError2()
}

func BenchmarkPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		panicFunc()
	}
}
