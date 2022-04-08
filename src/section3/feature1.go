//Go 특징(1)
package main

import "fmt"

func main() {
	//GO는 모호하거나 애매한 문법을 금지 (GO의 철학)
	//후치(후위) 연산자 허용 i++, 전치(전위) 연산자는 비허용 ++i -> 문법 오류발생
	//증감연산 반환 값 없음
	//포인터(Pointer 허용, 연산 비허용)
	//주석 (//, /**/) 사용법 숙지

	//ex1
	sum, i := 0, 0
	for i <= 100 {
		// sum += i++ 허용하지 않는다.
		sum += i
		i++
		// ++i 는 허용하지 않는다.
	}
	fmt.Println("ex1 : ", sum)
}
