package main

import (
	"fmt"
	"math"
)

func main() {
	// var num1 uint8 = math.MaxUint8 + 1 // 컴파일 에러 발생
	var num1 uint8 = math.MaxUint8

	fmt.Println(num1 + 1) // 이 값은 허용한다. 단 런타임 오류를 뱉지 않고 정상적으로 작동한다.

	// var num2 uint8 = 0 - 1; // 오버플로우 오류 발생 (언더플로우도 오버플로우로..)
}
