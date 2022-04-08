package main

import (
	"fmt"
	"sort"
)

func main() {
	// slice 추출 및 정렬
	// slice[i:j] -> i부터 j-1까지 추출
	// slice[i:] -> i부터 마지막 까지 추출
	// slice[:j] -> j부터 마지막 까지 추출
	// slice[:] -> 처음-부터 마지막까지 추출

	// ex1 일반추출
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ex1 : ", s1[:])
	// fmt.Println("ex1 : ", s1[0:])
	// fmt.Println("ex1 : ", s1[:5])
	// fmt.Println("ex1 : ", s1[0:len(s1)])
	// fmt.Println("ex1 : ", s1[3:])
	// fmt.Println("ex1 : ", s1[:3])
	// fmt.Println("ex1 : ", s1[1:3])

	// ex2 (sort package)

	s2 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}
	s3 := []string{"b", "d", "f", "a", "c", "e"}

	fmt.Println(s2)
	fmt.Println("ex2 : ", sort.IntsAreSorted(s2)) // 정렬확인
	sort.Ints(s2)                                 // 정렬 처리
	fmt.Println(s2)
	fmt.Println("ex2 : ", sort.IntsAreSorted(s2))

	fmt.Println()
	fmt.Println(s3)
	fmt.Println("ex3 : ", sort.StringsAreSorted(s3))
	sort.Strings(s3)
	fmt.Println(s3)
	fmt.Println("ex3 : ", sort.StringsAreSorted(s3))
}
