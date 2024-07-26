package main

import (
	"fmt"
	"math/big"
	mrand "math/rand"  // Alias to avoid conflict with crypto/rand
	"crypto/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4.5
    var first int = 6
	var second float64 = 6.8282929299292
	fmt.Println("The total is: ",second+ float64(first))
	fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// Generate a random number using math/rand
	mrand.Seed(time.Now().UnixNano())
	fmt.Println("Random number using math/rand:", mrand.Intn(5)+1)

	// Generate a random number using crypto/rand
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random number using crypto/rand:", myRandomNum)
}