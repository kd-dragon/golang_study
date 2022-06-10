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
	//파일 쓰기
	//Create : 새 파일 작성 및 파일 열기 (파일과 연결해서 파일객체를 리턴)
	//Close : 리소스 닫기
	//Write, WriteString, WriteAt : 파일 쓰기
	//각 운영체제 권한 주의(오류 메시지 확인)
	//예외 처리 정말 중요!

	//파일 쓰기 예제
	file, err := os.Create("D:/workspace-go/src/section11/test_write.txt")
	errCheck1(err)

	//리소스 해제
	defer file.Close()

	s1 := []byte{48, 49, 50, 51, 52}
	n1, err := file.Write([]byte(s1)) //문자열 -> byte 슬라이스 형으로 변환 후 쓰기
	errCheck2(err)

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes) \n", n1)

	file.Sync() //Write Commit 하는 것 ! stable 방식이므로 권장함.

	//쓰기 예제2
	s2 := "\nHello Golang! \n File Write Test! -1 \n"
	n2, err := file.WriteString(s2)
	errCheck2(err)

	fmt.Printf("쓰기 작업(2) 완료 (%d bytes) \n", n2)

	file.Sync()

	//ex3
	s3 := "Test WriteAt! -2 \n"
	n3, err := file.WriteAt([]byte(s3), 90) // len(offset) 조절하면서 테스트
	errCheck1(err)
	fmt.Printf("쓰기 작업(3) 완료 (%d bytes) \n", n3)

	//ex4
	n4, err := file.WriteString("Hello Golang! \n file Write Test! -3 \n")
	errCheck1(err)
	fmt.Printf("쓰기 작업(4) 완료 (%d bytes) \n", n4)
}
