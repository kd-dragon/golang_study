package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {

	//파일 읽기
	//csv 파일 읽기 예제

	//파일 생성
	file, err := os.Open("D:/workspace-go/src/section11/sample.csv")
	errCheck1(err)

	//리소스 해제
	defer file.Close()

	//CSV Reader 생성
	// rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) //권장

	//csv 내용 읽기
	row1, err1 := rr.Read() //1개의 Row 단위로 읽기
	errCheck1(err1)

	row2, err2 := rr.Read() //1개의 Row 단위로 읽기
	errCheck1(err2)

	fmt.Println("CSV Read Example")
	//fmt.Println(row)
	// fmt.Println(row)
	fmt.Println(row1[:])
	fmt.Println(row2[:])
	fmt.Println("==========================================")

	//CSV 내용 다가져오기
	rows, err := rr.ReadAll() //row 전체 읽기
	errCheck2(err)
	fmt.Println("CSV ReadAll Example")
	fmt.Println(rows)

	// Row 단위로 CSV 파일 읽기
	for _, row := range rows {
		//fmt.Println(i, row)
		for _, val := range row {
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}
}
