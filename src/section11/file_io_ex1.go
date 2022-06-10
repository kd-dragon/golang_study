package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {
	//파일 읽기, 쓰기 -> ioutil패키지 활용
	//아래 메소드 확인가능
	//WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	//https://golang.org/pkg/io/ioutil

	s := "Hello Golang!\nFile Write Test!\n"

	//파일모드(chmod, chown, chgrp) -> 퍼미션
	//읽기:4, 쓰기:2, 실행:1
	//소유자,그룹,기타사용자 순서

	err := ioutil.WriteFile("D:/workspace-go/src/section11/test_write1.txt", []byte(s), os.FileMode(0644)) //Go에서 퍼미션을 줄때는 앞에 0을 넣는다.
	errCheck(err)

	data, err := ioutil.ReadFile("D:/workspace-go/src/section11/sample.txt")
	errCheck(err)

	fmt.Println("===============================================")
	fmt.Println(string(data))
	fmt.Println("===============================================")
}
