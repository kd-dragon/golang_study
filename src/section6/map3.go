package main

import "fmt"

func main() {
	//Map 값 변경 및 삭제
	//ex1
	m1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com",
	}

	fmt.Println("ex1 : ", m1)
	m1["home2"] = "http://test2.com"
	fmt.Println("ex1 : ", m1)
	m1["home2"] = "http://home2.com"
	fmt.Println("ex1 : ", m1)

	//ex2 (delete)
	delete(m1, "home2")
	fmt.Println("ex1: ", m1)
	delete(m1, "google")
	fmt.Println("ex1: ", m1)
}
