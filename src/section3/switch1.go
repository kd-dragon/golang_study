package main

import "fmt"

func main() {
	// 조건문 - switch
	/*
			   Switch 뒤 표현식(Expression) 생략 가능
		     case 뒤 표현식(Expression) 사용 가능
		     자동 break 때문에 fallthrouth 존재
		     Type 분기 -> 값이 아닌 변수 Type으로 분기 가능
	*/

	//ex1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	//ex2 (go 에서 추천하는 패턴)
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	//ex3 문자사용 (c를 선언해서 case문에서 c == "go"라는 코드를 "go"로만 사용 가능)
	switch c := "go"; c {
	case "go":
		fmt.Println("GO!")
	case "java":
		fmt.Println("JAVA!")
	default:
		fmt.Println("일지하는 값 없음!")
	}

	//ex4 (짧은 선언 뒤 case와 비교할 값을 연산처리가능)
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("GO Lang!")
	case "java":
		fmt.Println("JAVA!")
	default:
		fmt.Println("일지하는 값 없음!")
	}

	// ex5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j와 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}
}
