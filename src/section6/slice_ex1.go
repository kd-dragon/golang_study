//자료형 : 슬라이스 심화(1)

package main

import "fmt"

func main() {

	// ex1
	// *주의: cap을 설정한 뒤에 동적으로 길이가 cap을 넘어가면 원래 cap의 기존의 2배 메모리를 할당한다. (성능 이슈 발생)
	s1 := []int{1, 2, 3, 4, 5} //len 5 cap 5
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)  // len cap dynamically increase
	s2 = append(s1, s2...) // slice를 삽입할 경우 ... 사용
	s3 = append(s2, s3...) // 추출 후 병합
	s3 = append(s3, s1...) // 추출 후 병합

	fmt.Printf("ex1 : %v %d \n", s1, cap(s1))
	fmt.Printf("ex1 : %v %d \n", s2, cap(s2))
	fmt.Printf("ex1 : %v %d \n", s3, cap(s3)) // cap이 직전 cap의 2배로 할당된다.

	// ex2
	s4 := make([]int, 0, 5)

	for i := 0; i < 15; {
		s4 = append(s4, i)
		i++
		fmt.Printf("ex2 -> len : %d, cap: %d, value: %v\n", len(s4), cap(s4), s4)
	}

}
