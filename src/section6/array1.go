package main

import "fmt"

func main() {
	/*배열
	  배열은 용량, 길이 항상 같다.
	  #배열 vs 슬라이스 차이점(중요)
	  1. 길이 고정 <-> 길이 가변
	  2. 값 타입 <-> 참조 타입
	  3. 복사 전달 <-> 참조 값 전달
	  4. 비교 연산자 사용 가능 <-> 비교 연산자 사용 불가
	  *결론 - 대부분 슬라이스를 사용한다.

	  *함수
	  cap(): 배열, 슬라이스 용량
	  len(): 배열, 슬라이스 갯수
	*/

	// ex1
	var arr1 [5]int                      //int형 기본값 0으로 초기화
	var arr2 [5]int = [5]int{1, 2, 3, 4} //[n]선언된 크기값은 일치해야한다. 원소는 n이하로만 선언되면 문제없음. (단, 0으로 초기화 됨)
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} // 배열크기 자동 맞춤 (거의 쓰이지 않는다. 참고만)
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, //마지막 콤마
	}
	arr1[2] = 5
	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T%d %v\n", arr7, len(arr7), arr7)
	fmt.Println(arr7)

	arr8 := [5]int{1, 2, 3, 4, 5}

	// *clean code
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "park"}
	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)

}
