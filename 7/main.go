//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mapa := map[int]int{}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(1)

	go func() {
		for i := 0; i < 5; i += 1 {
			mu.Lock()
			mapa[i] = i
			mu.Unlock()
			fmt.Println(i, "is written is map concurently")
			time.Sleep(1 * time.Second)
		}
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println(mapa)
}
