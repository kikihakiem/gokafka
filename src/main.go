package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"time"
	// "unicode/utf8"
)

func main() {
	r := gin.Default()
	r.GET("/*number", func(c *gin.Context) {
		param := strings.Trim(c.Param(("number")), "/")
		number, err := strconv.Atoi(param)

		if param != "" && err != nil {
			c.JSON(400, gin.H{
				"status":  "Error",
				"message": "Bad request: PATH should only contain integer or empty.",
			})

			return
		}
		if number == 0 {
			now := time.Now()
			number = now.Second()
		}

		ch := make(chan int) // Create a new channel.
		go Generate(ch)      // Launch Generate goroutine.
		var prime_nums []int

		for i := 0; i < number; i++ {
			prime := <-ch
			prime_nums = append(prime_nums, prime)
			ch1 := make(chan int)
			go Filter(ch, ch1, prime)
			ch = ch1
		}

		c.JSON(200, gin.H{
			"status":               "OK",
			"number":               number,
			"generated_prime_nums": prime_nums,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}
