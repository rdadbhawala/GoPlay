package errorLogging

import "testing"

// BenchmarkHandle ...
func BenchmarkHandleFF(b *testing.B) {
	runBenchmarkHandle(b, false, false)
}

func BenchmarkHandlePanicFF(b *testing.B) {
	runBenchmarkHandlePanic(b, false, false)
}

func BenchmarkHandleNoDefer(b *testing.B) {
	w := newWeb(newSrv(newDal(false, false)))
	for i := 0; i < b.N; i++ {
		w.HandleUnpanic()
	}
}

func BenchmarkHandleFT(b *testing.B) {
	runBenchmarkHandle(b, false, true)
}

func BenchmarkHandlePanicFT(b *testing.B) {
	runBenchmarkHandlePanic(b, false, true)
}

func BenchmarkHandleTF(b *testing.B) {
	runBenchmarkHandle(b, true, false)
}

func BenchmarkHandlePanicTF(b *testing.B) {
	runBenchmarkHandlePanic(b, true, false)
}

func runBenchmarkHandle(b *testing.B, one bool, two bool) {
	w := newWeb(newSrv(newDal(one, two)))
	for i := 0; i < b.N; i++ {
		w.Handle()
	}
}

func runBenchmarkHandlePanic(b *testing.B, one, two bool) {
	w := newWeb(newSrv(newDal(one, two)))
	for i := 0; i < b.N; i++ {
		w.HandlePanic()
	}
}
