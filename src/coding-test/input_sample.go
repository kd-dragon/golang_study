// GoLang 콘솔 입력 값 처리
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//엔터 기준으로 한줄씩 스캔 scn.Scan(), 스캔된 입력 값 scn.Text()
//띄어쓰기 기준으로 스캔 fmt.Scanln(&input)
func main() {
	scn := bufio.NewScanner(os.Stdin)
	var input string
	fmt.Scanln(&input)

	leng, _ := strconv.Atoi(input) // 입력받은 데이터 정수 변환
	var lines []string

	fmt.Println("len :: ", leng)

	for i := 0; i < leng; i++ {
		scn.Scan()                        // 엔터기준 한줄씩 입력 [(1 2 3), (4 5 6), (7 8 9)]
		lines = append(lines, scn.Text()) //데이터 배열로 추가
	}

	// input 4
	// input 1 2 3
	// input 4 5 6
	// input 7 8 9
	fmt.Println("lines :: ", lines)
	fmt.Println("len :: ", reflect.TypeOf((lines)))
	fmt.Println("idx 0 :: ", lines[0]) // 1 2 3
	fmt.Println("idx 1 :: ", lines[1]) // 4 5 6
}
