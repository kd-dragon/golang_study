package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	//외부 저장소 패키지 설치
	//2가지 방법
	//1. import 선언 후 폴더 이동후 go get 설치 -> 사용
	//2. go get 패키지 주소 설치 -> 선언

	// 선언 후 go get 설치 예제(엑셀 파일 읽기)
	xfile := "D:/workspace-go/src/section12/sample.xlsx"

	xlFile, err := xlsx.OpenFile(xfile)

	if err != nil {
		panic("Excel Loads Error!")
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				test := cell.String()
				fmt.Printf("%s\t", test)
			}
			fmt.Println()
		}
	}
}
