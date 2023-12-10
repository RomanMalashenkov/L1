/*
Написать программу, которая конкурентно рассчитает значения квадратов
чисел, взятых из массива (2,4,6,8,10), и выведет их квадраты в Stdout
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func square(n int, mu *sync.Mutex) {
	mu.Lock()
	fmt.Println(n * n)
	mu.Unlock()
}

func main() {
	bas := [5]int{2, 4, 6, 8, 10}
	var mu sync.Mutex

	for _, v := range bas {
		go square(v, &mu)
	}

	time.Sleep(1 * time.Millisecond)
}
