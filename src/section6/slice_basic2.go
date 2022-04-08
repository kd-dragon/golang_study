package main

import "fmt"

func main() {
	// 슬라이스 참조 타입 증명

	// ex1 (배열 복사 증명)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int
	arr2 = arr1
	arr2[0] = 7

	fmt.Println(arr1)
	fmt.Println(arr2)

	// ex2 (슬라이스 참조 증명)
	slice1 := []int{1, 2, 3}
	var slice2 []int
	slice2 = slice1
	slice2[0] = 7
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Printf("ex2 : %p %v\n", &slice1, slice1)
	fmt.Printf("ex2 : %p %v\n", &slice2, slice2)

	// ex3 (slice 예외)
	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4]) //default value
	// fmt.Println("ex3 : ", slice3[50]) //길이만큼 초기화 되기때문에 50을 출력시 index out of range 예외 발생

	// ex4

	for i, v := range slice1 {
		fmt.Println("ex4 : ", i, v)
	}
}
