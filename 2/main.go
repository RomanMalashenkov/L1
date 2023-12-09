/*
Написать программу, которая конкурентно рассчитает значения квадратов
чисел, взятых из массива (2,4,6,8,10), и выведет их квадраты в Stdout
*/

package main

import (
	"fmt"
	"sync"
)

func square(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	n *= n
	fmt.Println(n)
}

func main() {
	bas := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(bas))

	for _, v := range bas {
		go square(v, &wg)
	}

	wg.Wait()
}
