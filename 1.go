//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской
//структуры Human (аналог наследования).

package t1

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) Person() {
	fmt.Printf("%s, %d\n%s", h.name, h.age, h.phone)
}

type Action struct {
	stadt string
	Human // встроили структуру Human в структуру Action
}

func t1() {

	a := Action{
		"Bryansk",
		Human{"Roman", 23, "+79123456789"},
	}

	//вызываем метод, который за счет встраивания структур может использоваться
	//другой структурой, в которую мы встроили данную
	a.Person()
}
