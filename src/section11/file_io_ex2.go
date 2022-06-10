package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {

	//ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현함 -> 즉, Writer, Read 메소드 사용 가능
	/*
	   type Reader interface {
	     Read(p []byte) (n int, err error)
	   }
	   type Writer interface {
	     Write(p []byte) (n int, err error)
	   }
	*/

	//즉, bufio의 NewReader, NewWriter를 통해서 객체 반환 후 메소드 사용
	//bufio(Buffered io) 패키지

	//https://golang.org/pkg/bufio

	//파일 열기
	//두번째 매개변수 확인
	//https://golang.org/pkg/os/#pkg-constants

	/*
	   상태 (4byte 버퍼일경우)
	   a -----> a
	   b -----> ab
	   c -----> abc
	   d -----> abcd
	   e -----> e      ------> 텍스트파일에는 abcd가 써진다.
	   f -----> ef
	   g -----> efg
	   h -----> efgh
	   i -----> i      ------> abcdefgh
	*/

	//os.O_CREATE|os.O_RDWR ---> 없으면 만들거나 있으면 읽고 쓰고 할게 있게 하라는 flag
	file, err := os.OpenFile("D:/workspace-go/src/section11/test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	//bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) //Writer 반환(버퍼 사용)
	wt.WriteString("Hello Golang!\n")
	wt.Write([]byte("Hello Golang2!\n"))
	//에러체크
	errCheck(err)

	//버퍼 정보 출력하기
	fmt.Printf("사용한 Buffer Size (%d byte) \n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d byte) \n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d byte) \n", wt.Size())

	wt.Flush() //*버퍼를 비우고 반영하는 것! (버퍼의 내용을 디스크에 기록)

	fmt.Println("쓰기 작업 완료")
	fmt.Println("===========================================================================")

	rt := bufio.NewReader(file) //Reader 반환
	f1, err := file.Stat()
	errCheck(err)

	b := make([]byte, f1.Size())

	fmt.Println("파일 정보 출력 : ", f1)
	fmt.Println("파일 명 : ", f1.Name())
	fmt.Println("파일 크기 : ", f1.Size())
	fmt.Println("파일 수정 시간 : ", f1.ModTime())

	fmt.Println("===========================================================================")

	file.Seek(0, os.SEEK_SET)

	data, _ := rt.Read(b) //읽기(ReadLine, ReadByte, ReadBytes 등)
	//rt.Read(b)

	fmt.Printf("전체 Buffer Size : (%d bytes) \n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("===========================================================================")
	fmt.Println(string(b))
	fmt.Println("===========================================================================")

	defer file.Close()
}
