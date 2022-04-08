//데이터 타입: Numeric 연산 (1)
package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자 연산 (산술, 비교)
	//타입이 같아야 가능 (엄격)
	//다른 타입끼리는 반드시 형 변환 후 연산
	//형 변환 없을 경우 예외(에러) 발생
	// 종류: +, -, *, %, /, <<, >>, &, ^

	//예제1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex1 : ", n1)
	fmt.Println("ex1 : ", n2)
	fmt.Println("ex1 : ", n3)
	fmt.Println("ex1 : ", n4)
	fmt.Println("ex1 : ", math.MaxInt8)
	fmt.Println("ex1 : ", math.MaxInt16)
	fmt.Println("ex1 : ", math.MaxInt32)
	fmt.Println("ex1 : ", math.MaxInt64)

	n5 := 10000 //int
	n6 := int16(10000)
	n7 := uint8(100)

	//fmt.Println("ex2 : ", n5+n6) mismatched type 에러 발생
	fmt.Println("ex2 : ", n5+int(n6))
	fmt.Println("ex2 : ", n6+int16(n7))

	//비교 연산자도 타입이 같아야 비교가능
	fmt.Println("ex2 : ", n6 > int16(n7))
}
