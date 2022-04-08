package main

import (
	"fmt"
)

func main() {
	// MAP 조회 및 순회 (Iterator)

	m1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex1 : ", m1["google"])
	fmt.Println("ex1 : ", m1["daum"])
	fmt.Println()

	//ex2 (순서 없으므로 랜덤)
	for k, v := range m1 {
		fmt.Println("ex2 : ", k, v)
	}
	fmt.Println()
	for _, v := range m1 {
		fmt.Println("ex2 : ", v)
	}

}
