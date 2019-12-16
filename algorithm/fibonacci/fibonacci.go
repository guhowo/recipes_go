package fibonacci

func Fib(n int) int {
	f1, f2 := 0, 1
	for i := 2; i < n; i++ {
		f1 = f1 + f2
		f2 = f1 + f2
	}

	return f2
}
