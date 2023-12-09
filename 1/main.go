/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) Person() {
	fmt.Printf("%s, %d\nPhone: %s", h.name, h.age, h.phone)
}

type Action struct {
	stadt string
	Human // встроили структуру Human в структуру Action
}

func main() {

	a := Action{
		"Bryansk",
		Human{"Roman", 23, "+79123456789"},
	}

	a.Person()
}
