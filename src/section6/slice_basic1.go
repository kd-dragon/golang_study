package main

import (
	"fmt"
)

func main() {
	/* 슬라이스
	길이 고정 X (가변) -> 동적으로 크기가 늘어남
	레퍼런스(참조) 타입
	슬라이스 (길이 & 용량) 크기 동적으로 할당 가능

	# 2가지 선언 방법
	1. 배열처럼 선언
	2. make 함수: (자료형, 길이, 용량(생략시 길이값) 설정)
	*/

	// ex1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{}

	// 완벽하게 초기화되지 않은 뒤 값을 수정하면 index out of range 오류 발생
	// slice2[0] = 1
	slice3[4] = 10 //값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)
	fmt.Println("================================================================")

	// ex2 (make 함수 사용 -> 자동 초기화 해준다.)
	var slice5 []int = make([]int, 5, 10) // 길이 5, 용량 10인 int형 슬라이스 ( 용량 10의 메모리를 할당해놓고 5만큼 사용하는 것 )
	var slice6 = make([]int, 5, 5)        // 가장 많이 사용하는 패턴
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7 //삽입

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	// ex3
	var slice9 []int //nil 슬라이스(길이와 용량 0)

	if slice9 == nil {
		fmt.Println("This is Nil")
	}

}
