// 1이 될때까지
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

	*/

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, K int
	var result int
	fmt.Fscan(reader, &N, &K)

	t1 := time.Now()
	for N != 1 {
		if N%K != 0 {
			result++
			N--
		} else {
			result++
			N /= K
		}
	}
	fmt.Println("My Answer :: ", time.Since(t1), N, K)
	fmt.Fprintln(writer, result)
}
