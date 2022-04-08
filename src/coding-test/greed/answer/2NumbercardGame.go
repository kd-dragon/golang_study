// 숫자 카드 게임
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 가장 높은 숫자가 쓰인 카드 한 장 뽑기 (MAX)
	/*
		   # 조건
		   1. 카드는 N x M 행렬 형태 (2중 배열)
		   2. 우선적으로 행 선택, 그 행에서 가장 낮은 숫자 카드 뽑기
		   3. 각 행마다 낮은 카드 중에 가장 높은 카드를 뽑기
			 4. 각 숫자는 최소 1 최대 10000의 자연수
	*/
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var max int
	for i := 0; i < n; i++ {
		min := 10001
		fmt.Print("- ")
		for _, item := range make([]int, m) {
			fmt.Fscan(reader, &item)
			if min > item {
				min = item
			}
		}
		if max < min {
			max = min
		}
	}
	fmt.Fprintln(writer, "max : ", max)
}
