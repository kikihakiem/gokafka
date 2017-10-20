package controllers

import (
	// "log"
	"github.com/gin-gonic/gin"
	"my_helper"
	"time"
	// "unicode/utf8"
)

func Prime(c *gin.Context) {
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
}
