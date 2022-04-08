package main

import "fmt"

func main() {
	//ex1 break문 사용법 - break시 Label로 이동가능 (자주 사용되지는 않지만 알아두기)
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			//fmt.Println("ex1 : ", j)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break
			}
			//fmt.Println("ex1-2 : ", j)
		}
	}

	//------------------------------------

	//ex2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2 : ", i)
	}

	// ex3 label 활용 (Label 뒤에는 반드시 for문와야함)
Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3 : ", i, j)
		}
	}
}
