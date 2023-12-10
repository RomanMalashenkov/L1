// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
// func main() {
// 	someFunc()
// }

/*
Ответ:
Здесь происходит утечка памяти, когда someFunc отрабатывает, то в justString, хранятся 100
символов большой строки v, память под этот новый срез не выделяется, и мы имеем доступ только
к этим 100 символам огромной строки, а оставшаяся часть висит в памяти и не может быть
очищена сборщиком мусора, так как JustString ссылается на кусок этой памяти. Кроме того,
JustString - глобальная переменная, использование которых может привести к неконтролируемым
последствиям (например, ошибки, которые трудно отладить)
*/
package main

import (
	"fmt"
	"strings"
)

func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	// Копируем строку, выделив для копии новый участок памяти
	*justString = strings.Clone(v)[:100]
}

func createHugeString(length int) string {
	// Заполняем строку, пока она не станет нужной длины
	res := ""
	for length > 0 {
		res += "test"
		length -= 4
	}
	return res
}

func main() {
	var justString string
	someFunc(&justString)
	fmt.Println(justString)
}
