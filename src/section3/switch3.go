package main

import "fmt"

func main() {
	//ex1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // a가 2 또는 4 또는 6인 경우 (AND)
		fmt.Println("a -> ", a, " 짝수이다.")
	case 1, 3, 5: // a가 1,3,5 인 경우
		fmt.Println("a -> ", a, " 짝수이다.")
	}

	// ex2 (fallthrouth) : 조건에 맞는 case의 다음 case문을 실행시키는 명령어
	switch e := "go"; e {
	case "java":
		fmt.Println("java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
		fallthrough
	case "ruby":
		fmt.Println("ruby!")
		fallthrough
	case "c":
		fmt.Println("C!")
		//fallthrough (오류발생: cannot fallthrough final case in switch)
	}

}
