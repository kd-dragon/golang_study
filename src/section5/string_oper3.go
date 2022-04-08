package main

import (
	"fmt"
	"strings"
)

func main() {
	//ex1 (결합 : 일반연산)
	/*
	   string은 한번 선언하면 수정이 불가하다. (Immutable) 따라서 매번 새롭게 생성된다.
	*/
	str1 := "The Go programming language is an open source project to make programmers more productive." +
		"Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked"

	str2 := "Introduces the basics of creating and using multi-module workspaces in Go. Multi-module workspaces are useful for making changes across multiple modules."

	fmt.Println("ex1 : ", str1+str2)

	//ex2 (결합 : Join) -> 더 효율적인 성능
	strSet := []string{} // 슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex2 : ", strings.Join(strSet, "-----"))
	fmt.Println("ex2 : ", strings.Join(strSet, ""))
}
