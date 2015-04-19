// fib_test project fib_test.go
package fib_test

import (
	"fibutil/fib"
	"testing"
)

func TestFib(t *testing.T) {
	ar := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}

	nextFibo := fib.Fibo()
	for i := 0; i < len(ar); i++ {
		j := nextFibo()
		if ar[i] != j {
			t.Fatalf("%d, %d!=%d", i, ar[i], nextFibo())
		}
	}
}

func BenchmarkFib(b *testing.B) {
	nextFibo := fib.Fibo()
	for i := 0; i < b.N; i++ {
		nextFibo()
	}
}
