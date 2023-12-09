/*
Дана последовательность чисел: 2, 4, 6, 8, 10.
Найти сумму их квадратов с использованием конкурентных вычислений
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	bas := [5]int{2, 4, 6, 8, 10}
	sl := make([]int, 0)
	res := 0

	var mu sync.Mutex

	for _, v := range bas {
		go func(n int) {
			mu.Lock()
			m := n * n
			sl = append(sl, m)
			mu.Unlock()
		}(v)
	}

	time.Sleep(10 * time.Millisecond)

	for _, v := range sl {
		res += v
	}

	fmt.Println(res)
}
