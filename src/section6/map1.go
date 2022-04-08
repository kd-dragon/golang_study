package main

import (
	"fmt"
)

func main() {
	/* map[type]type
	Hashtable, Dictionary(python), key-value data 자료 구조
	Reference Type (mutable)
	비교 연산자 사용 불가 (참조 타입이므로)

	# 특징
	참조타입을 key로 사용 불가능, 값(value) 으로 모든 타입 사용 가능
	make 함수 및 축약(리터럴)로 초기화 가능
	순서 없음 O(1) 복잡도
	*/

	// ex1 map[key-type] value-type
	var m1 map[string]int = make(map[string]int) // 정석 사용 패턴
	var m2 = make(map[string]int)                // 자료형 생략 패턴
	m3 := make(map[string]int)                   // 주로 사용하는 패턴(짧은 선언)

	fmt.Println("ex1 : ", m1)
	fmt.Println("ex1 : ", m2)
	fmt.Println("ex1 : ", m3)

	// ex2
	m4 := map[string]int{} // 형태만 JSON형태로 선언 -> Json a = {"a":1, "b":"hi"}
	m4["apple"] = 25
	m4["banana"] = 40
	m4["orange"] = 33

	m5 := map[string]int{ // 선호하는 방식
		"apple":  15,
		"banana": 40,
		"orange": 33,
	}

	m6 := make(map[string]int)
	m6["apple"] = 25
	m6["banana"] = 40
	m6["orange"] = 33

	fmt.Println("ex2 m4 : ", m4)
	fmt.Println("ex2 m5 : ", m5)
	fmt.Println("ex2 m6 : ", m6)
	fmt.Println("ex2 : ", m6["orange"])
	fmt.Println("ex2 : ", m6["banana"])

}
