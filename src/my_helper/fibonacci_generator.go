package my_helper

func FibonacciGenerator(number int) (fibSequence []int) {
	f := fibonacci()
	for i := 0; i < number; i++ {
		fibSequence = append(fibSequence, f())
	}

	return
}

func fibonacci() func() int {
	prev, curr, next := 0, 0, 1

	return func() int {
		prev = curr
		curr = next
		next = prev + curr

		return prev
	}
}
