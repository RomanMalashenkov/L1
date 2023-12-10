// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

package main

import (
	"fmt"
	"time"
)

func input(n *int) error {
	fmt.Print("Enter working time in seconds: ")
	_, err := fmt.Scan(n)
	fmt.Println()
	return err
}

func main() {
	var n int

	for input(&n) != nil {
	}

	start := time.Now()

	ch := make(chan int)
	go func() {
		for {
			fmt.Println(<-ch)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	series := 1

	for ; ; series++ {
		if time.Now().Sub(start) >= time.Duration(n)*time.Second {
			return
		}
		ch <- series
	}
}
