// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	// с поимощью контекста
	ctx, stop := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 1 is killed")
				return
			default:
				fmt.Println("Goroutine 1 is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	stop()

	//через канал
	ch := make(chan int)

	go func() {
		for {
			select {
			case <-ch:
				close(ch)
				return
			default:
				fmt.Println("goroutine 1 is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	ch <- 1

	//через waitgroup
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for i := 0; i < 3; i += 1 {
			fmt.Printf("i am i: %d\n", i)
			time.Sleep(1 * time.Second)
		}
		defer wg.Done()
	}()

	wg.Wait()
}
