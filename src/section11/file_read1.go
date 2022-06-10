package main

import (
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
	//file read
	//Open : 기존 파일 열기
	//Close : 리소스 닫기
	//Read, ReadAt 파일 읽기
	//각 운영체제 권한 주의(오류 메시지 확인)

	//예외 처리 정말 중요!
	//탐색 Seek 중요
	//https://golang.org/pkg/os

	//파일 읽기 예제
	//파일 열기
	file, err := os.Open("D:/workspace-go/src/section11/sample.txt")
	errCheck1(err)

	//읽기 예제1
	fileInfo, err := file.Stat() //파일 사이즈 확인 위해 정보 획득
	errCheck2(err)

	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) : ", fileInfo)
	fmt.Println("파일명 (2) : ", fileInfo.Name())
	fmt.Println("파일크기 (3) : ", fileInfo.Size())
	fmt.Println("파일 마지막 수정시각 (4) : ", fileInfo.ModTime())
	fmt.Printf("읽기작업(1) 완료 (%d byte)\n\n", ct1)
	fmt.Println(fd1)
	fmt.Println(string(fd1))

	fmt.Println("=====================================================")

	//읽기 예제2(탐색 : Seek(Offset))
	o1, err := file.Seek(20, 0) // 0:처음 위치, 1:현재위치, 2:마지막위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)

	fmt.Printf("읽기작업(2) 완료 (%d byte) (%d ret)\n\n", ct2, o1) //o1 -> offset
	fmt.Println(string(fd2), "\n")
	fmt.Println("=====================================================")

	//읽기 예제3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) //offset 위치 부터 읽어온다.
	errCheck1(err)
	fmt.Printf("읽기작업(3) 완료 (%d byte) (%d ret)\n\n", ct3, o2) //o1 -> offset
	fmt.Println(string(fd3), "\n")
	fmt.Println("=====================================================")

	defer file.Close()
}
