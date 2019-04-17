package main

import (
	"fmt"
	"sync"

	color "github.com/fatih/color"
)

func main() {
	cyan := color.New(color.FgCyan)
	magenta := color.New(color.FgMagenta)

	up := func() {
		for i := 0; i < 100; i++ {
			cyan.Printf("U%d ", i)
		}
	}

	down := func() {
		for i := 100; i > 0; i-- {
			magenta.Printf("D%d ", i)
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		up()
		wg.Done()
	}()

	go func() {
		down()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println()
}
