package main

import "fmt"

func main() {
	//배열 순회

	// ex1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	// len 길이 반복
	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex1 : ", arr1[i])
	}
	fmt.Println()

	// ex2 *실제 가장 많이 사용되는 방법
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for _, v := range arr2 {
		fmt.Println("ex2 : ", v)
	}
	fmt.Println()
	// ex3 (인덱스만 사용하기)
	for v := range arr2 {
		fmt.Println("ex2 index : ", v)
	}

}
