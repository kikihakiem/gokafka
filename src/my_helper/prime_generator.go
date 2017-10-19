package my_helper

func GeneratePrimeNumbers(number int) (primeNums []int) {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Launch Generate goroutine.

	for i := 0; i < number; i++ {
		prime := <-ch
		primeNums = append(primeNums, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	return
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}
