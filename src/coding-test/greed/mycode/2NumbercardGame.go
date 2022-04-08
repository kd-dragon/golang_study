// 숫자 카드 게임
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// 가장 높은 숫자가 쓰인 카드 한 장 뽑기 (MAX)
	/*
	   # 조건
	   1. 카드는 N x M 행렬 형태 (2중 배열)
	   2. 우선적으로 행 선택, 그 행에서 가장 낮은 숫자 카드 뽑기
	   3. 각 행마다 낮은 카드 중에 가장 높은 카드를 뽑기
	*/
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	cards := make([][]int, n)

	for i := 0; i < n; i++ {
		cards[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &cards[i][j])
		}
	}

	fmt.Fprintln(writer, cards)

	var max int
	for _, item := range cards {
		sort.Ints(item)
		if max < item[0] {
			max = item[0]
		}
	}

	fmt.Fprintln(writer, "max : ", max)
}
