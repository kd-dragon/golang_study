package main

import (
	"fmt"
)

// init 메서드 중복 가능

func init() {
	fmt.Println("Init1 Method Start ! ")
}

func init() {
	fmt.Println("Init2 Method Start ! ")
}

func init() {
	fmt.Println("Init3 Method Start ! ")
}

func init() {
	fmt.Println("Init4 Method Start ! ")
}

func main() {
	fmt.Println("Main Method Start ! ")
}
