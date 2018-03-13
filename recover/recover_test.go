package recover

import "testing"

func BenchmarkRecover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recoverFunc(i)
	}
}

func recoverFunc(x int) int {
	defer func() {
		recover()
	}()
	return 100 - x
}

func BenchmarkSameFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sameFunc(i)
	}
}

func sameFunc(x int) int {
	defer func() {
	}()
	return 100 - x
}
