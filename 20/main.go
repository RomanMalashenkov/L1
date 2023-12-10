//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {

	//создаем новую строку
	base := "snow dog sun"
	fmt.Println(base)
	//создаем массив слов из строки, разделенной пробелами, и объявляем новую строку
	splited := strings.Split(base, " ")
	newString := ""

	for i := len(splited) - 1; i >= 0; i -= 1 {
		newString += splited[i] + " "
	}

	base = strings.TrimRight(newString, " ")
	fmt.Println(base)
}
