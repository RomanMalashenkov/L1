// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

package main

import "fmt"

type Set[V comparable] map[V]struct{}

func NewSetWithValue[V comparable](value ...V) Set[V] {
	set := make(Set[V], len(value))

	for _, v := range value {
		set[v] = struct{}{}
	}

	return set
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewSetWithValue(seq...)

	for v := range set {
		fmt.Println(v)
	}
}
