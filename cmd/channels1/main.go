// SOLUTION to channels exercise described below.

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

func main() {

	// Create an unbuffered channel.
	numChan := make(chan int)
	done := make(chan bool)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("R1", numChan)
		done <- true
	}()

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("R2", numChan)
		done <- true
	}()

	// Send a value to start the counting.
	numChan <- 0

	// Wait for the program to finish.
	<-done
}

// goroutine simulates sharing a value.
func goroutine(name string, c chan int) {
	for {

		// Wait for the value to be sent.
		// If the channel was closed, return.
		n, open := <-c
		if !open {
			return
		}

		// Display the value.
		fmt.Printf("%s: %2d\n", name, n)

		// Terminate when the value is 10.
		if n == 10 {
			return
		}

		// Increment the value and send it
		// over the channel.
		n++
		c <- n
	}
}
