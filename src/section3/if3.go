package main

import "fmt"

func main() {
	// 제어문 if elseif else

	i := 4

	//ex1
	if i >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120미만")
	} else if i < 100 && i >= 50 {
		fmt.Println("50 이상 100미만")
	} else {
		fmt.Println("50 미만")
	}

}
