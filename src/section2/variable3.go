//변수3
package main

import "fmt"

func main() {
	//짧은 선언 (GO에만 있는 문법)
	//함수 안에서만 사용 (전역변수로 사용불가), 선언 후 재할당시 오류 발생
	//일회성 사용, 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있다. -> 추천
	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// shortVar1 := 10 //재할당시 오류 발생
	fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

	// Example
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success! :", i)
	}

}
