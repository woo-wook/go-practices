package main

import (
	"fmt"
	"unicode"
)

func main() {
	var r1 = '한'
	fmt.Println(unicode.Is(unicode.Hangul, r1))
	fmt.Println(unicode.Is(unicode.Latin, r1))

	var r2 = '柳'
	fmt.Println(unicode.Is(unicode.Han, r2))
	fmt.Println(unicode.Is(unicode.Latin, r2))

	var r3 = 'a'
	fmt.Println(unicode.Is(unicode.Han, r3))
	fmt.Println(unicode.Is(unicode.Latin, r3))

	fmt.Println(unicode.IsGraphic('1'))
	fmt.Println(unicode.IsGraphic('a'))
	fmt.Println(unicode.IsGraphic('柳'))
	fmt.Println(unicode.IsGraphic('\n')) // display 안되는 문자로 false

	fmt.Println(unicode.IsLetter('a')) // true
	fmt.Println(unicode.IsLetter('1')) // false (숫자이기에)

	fmt.Println(unicode.IsDigit('1')) // true
	fmt.Println(unicode.IsDigit('a')) // false (문자)

	fmt.Println(unicode.IsControl('\n')) // 제어 문자이기에 true
	fmt.Println(unicode.IsControl('1'))  // 제어 문자이기에 true

	fmt.Println(unicode.IsMark('\u17c9')) // mark

	fmt.Println(unicode.IsSymbol('😋')) // 심볼
}
