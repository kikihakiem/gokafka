package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	"my_helper"
	"time"
	// "unicode/utf8"
)

func main() {
	r := gin.Default()
	r.GET("/prime/*number", func(c *gin.Context) {
		raw := c.Param("number")
		number, err := my_helper.GetNumber(&raw)

		if err != nil {
			if err.Type() == my_helper.NonIntegerStr {
				c.JSON(400, gin.H{
					"status":  "Error",
					"message": "Bad request: PATH should only contain integer or empty.",
				})

				return
			}

			// if no number specified, give random number of primes
			if err.Type() == my_helper.EmptyStr {
				number = time.Now().Second()
			}
		}

		primeNums := my_helper.GeneratePrimeNumbers(number)

		c.JSON(200, gin.H{
			"status":               "OK",
			"number":               number,
			"generated_prime_nums": primeNums,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
