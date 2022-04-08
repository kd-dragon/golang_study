package main

import (
	"fmt"
	"sort"
)

func main() {

	var n, m, k int
	fmt.Scanf("%d %d %d\n", &n, &m, &k)

	var slice []int = make([]int, n)

	for i := 0; i < n; {
		fmt.Scanf("%d", &slice[i])
		i++
	}
	sort.Ints(slice)
	fmt.Println("max, sub :", slice[n-1], slice[n-2])

	var output int

	// *Idea (수열 개념) k+1만큼의 크기가 반복적으로 더해진다.
	count := m / (k + 1) * k // count => max 값이 곱해질 총 횟수
	count += m % (k + 1)     //(k+1) 나눈 나머지는 모두 max값이 더해진다.

	output = count * slice[n-1]        // 최대값 * count
	output += (m - count) * slice[n-2] // 전체 m에서 count를 빼면 두번째 큰 수가 더 해지는 횟수

	fmt.Println("final value ::: ", output)
}
