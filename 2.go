/*
Написать программу, которая конкурентно рассчитает значения квадратов
чисел, взятых из массива (2,4,6,8,10), и выведет их квадраты в Stdout
*/

package t2

import (
	"fmt"
	"time"
)

func t2() {
	bas := [5]int{2, 4, 6, 8, 10}

	for _, v := range bas {
		go func(n int) {
			n *= n
			fmt.Println(n)
		}(v)
	}
	time.Sleep(30 * time.Millisecond)
}
