// 1이 될때까지 (교재 정답)
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	/*
		   임의의 수 N (2 <= N <= 100,000)이 1이 될 때까지 다음 조건 수행
		   1. N에서 1을 뺀다.
		   2. N을 K로 나눈다. 단, 나누어떨어질때만 선택 가능 K (2 <= K <= 100,000)

		   ex) N = 17, K = 4
		   1) N = N - 1 (:=16)
		   2) N = N / K (:=4)
		   3) N = N / K (:=1)

		   핵심 원리: 먼저 조건1로 4의 배수를 만든다. 1이 될때까지 나눈다.

			 성능 개선 고려사항
			  - 증감을 반복문을 통하지 않고 한번에 할 수 있는 수학적인 연산을 이용하자
				- ex) 3 + 3 + 3 = 3*3 <-> for i:=0 i<3 i++ { cnt+=3 } 자원낭비...
	*/

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, K int
	var result int
	fmt.Fscan(reader, &N, &K)

	t1 := time.Now()
	var target int
	for {

		// 나머지를 구하는 식 (N - (N/K) * K) ==> 27 - 27/4 * 4 = 27 - 24 = 3
		target = (N / K) * K
		result += (N - target)
		N = target

		// N이 K보다 작을 때 반복문 탈출
		if N < K {
			break
		}

		// K로 나누기
		result++
		N /= K
	}
	// 마지막으로 남은 수에 대하여 1씩 빼기
	result += (N - 1) // 1씩 뺴는 것을 한번에 빼준다. N-1만큼 빼줘야함
	fmt.Println("book Answer :: ", time.Since(t1), N, K)

	fmt.Fprintln(writer, result)
}
