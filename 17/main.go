// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
)

func binarySearch(arr []int, elem int) int {

	f := 0
	l := len(arr) - 1

	for {
		mid := (f + l) / 2
		if arr[mid] == elem {
			return mid
		} else if elem < arr[mid] {
			l = mid - 1
		} else {
			f = mid + 1
		}

		if f < 0 || l >= len(arr) {
			return -1
		}
	}
}

func main() {
	arr := []int{1, 3, 6, 7, 8, 9, 16, 77}
	item := arr[5]
	fmt.Println(arr)
	fmt.Printf("Item = %d, index = %d", item, binarySearch(arr, 9))
}
