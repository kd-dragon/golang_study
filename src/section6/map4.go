package main

import (
	"fmt"
)

func main() {
	// Map 조회할 경우 주의 할 점
	// 여러가지 조회 패턴
	m1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	v1 := m1["lemon"]
	v2 := m1["kiwi"]
	v3, ok := m1["kiwi"]

	fmt.Println("ex1 : ", v1)
	fmt.Println("ex1 : ", v2)
	fmt.Println("ex1 : ", v3, ok) // 두번째 리턴 값으로 키 존재 유무 확인

	// ex2 (만약 map에 key 'kiwi' 의 value 값이 있을 경우)
	if val, ok := m1["kiwi"]; ok {
		fmt.Println("ex2 : ", val)
	} else {
		fmt.Println("ex2 : kiwi is not exists")
	}

	if val, ok := m1["banana"]; ok {
		fmt.Println("ex2 : ", val)
	} else {
		fmt.Println("ex2 : banana is not exists")
	}
	// 반대로 없다면 !ok
	if _, ok := m1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exists!")
	}

}
