package errorLogging

func newWeb(s srv) web {
	return &webImpl{s}
}

type webImpl struct {
	s srv
}

func (w *webImpl) Handle() {
	err := w.s.Feature1()
	if err != nil {

	}
}

func (w *webImpl) HandleUnpanic() {
	w.s.Feature1Panic()
}

func (w *webImpl) HandlePanic() {
	defer w.handleRecover()
	w.s.Feature1Panic()
}

func (w *webImpl) handleRecover() {
	recover()
}
