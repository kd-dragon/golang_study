package main

import (
	"fmt"
)

func main() {
	//채널(Channel)
	//Close: 채널 닫기

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good!"
		}
	}()

	val1, ok1 := <-ch // ok1은 성공여부의 bool값
	fmt.Println("ex : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex : ", val3, ok3)

	close(ch)
	val4, ok4 := <-ch //ok4 값 false -> Close됬으므로
	fmt.Println("ex : ", val4, ok4)

}
