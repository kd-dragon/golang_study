package main

import "fmt"

func main() {
	//ex1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1 // sum = sum + 1
	}
	fmt.Println("ex1 : ", sum1)

	//ex2 (i++ inner)
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++ //Go에서 후치(후처리) 연산은 반환값이 없다. 따라서 j := i++ 사용 불가
	}
	fmt.Println("ex2 : ", sum2)

	//ex3 (while)
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)

	//ex4 (자주쓰이진 않음 - 이해만 하자)
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}

	//error case
	//ex1 (후치연산은 반환값이 없으므로 i++, j+=10에서 오류 발생)
	/*
	   for i, j := 0, 0; i <= 10; i++, j += 10 {
	     fmt.Println("ex4 : ", i, j)
	   }
	*/
}
