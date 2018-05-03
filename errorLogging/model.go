package errorLogging

type dal interface {
	Op1() error
	Op1Panic()

	Op2() error
	Op2Panic()
}

type srv interface {
	Feature1() error
	Feature1Panic()

	// Feature2() error
	// Feature2Panic()
}

type web interface {
	Handle()
	HandlePanic()
	HandleUnpanic()
}
