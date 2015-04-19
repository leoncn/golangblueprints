// fib project fib.go
package fib

func Fibo() func() int {
	n2, n1, i := 0, 0, 0

	return func() int {
		if i == 0 {
			n2, n1 = 0, 0
		} else if i < 3 {
			n2, n1 = 1, 1
		} else {
			n1, n2 = n2, n1+n2
		}
		i++
		return n2
	}
}
