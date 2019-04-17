// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Add imports.

// Declare constant for number of goroutines.
const NRoutines = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffered channel with room for
	// each goroutine to be created.
	nums := make(chan int, NRoutines)

	// Iterate and launch each goroutine.
	for i := 0; i < NRoutines; i++ {

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			nums <- rand.Intn(7777)
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	msgsReceived := NRoutines

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var myInts []int
	for msgsReceived > 0 {
		myInts = append(myInts, <-nums)
		msgsReceived--
	}
	close(nums)

	// Print the values in our slice.
	fmt.Println(myInts)
}
