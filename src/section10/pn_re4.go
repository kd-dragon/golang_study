package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	//defer 함수
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("file open error : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1 : ", f.Name())
	}

	defer f.Close() //리소스 닫기 중요!
}

func main() {
	fmt.Println("Main Start")
	fileOpen("undefined.txt")
	fmt.Println("Main End!")
}
