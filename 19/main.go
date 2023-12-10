// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func reverseString(str string) string {

	strRune := []rune(str)

	for pos, _ := range strRune {
		buf := strRune[pos]
		strRune[pos] = strRune[len(strRune)-pos-1]
		strRune[len(strRune)-pos-1] = buf

		if pos == (len(strRune)-1)/2 {
			break
		}
	}
	return string(strRune)
}

func main() {
	fmt.Print("Enter string: ")
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	fmt.Printf("Reversed: %s", reverseString(s))

}
