package main

import "fmt"

func main() {
	//구조체 익명 선언 및 활용

	// ex1 type 구조체명 타입
	car1 := struct{ name, color string }{"520d", "red"} //익명구조체 (이렇게 많이 씀)
	fmt.Println("ex1 : ", car1)
	fmt.Printf("ex1 : %#v\n", car1)

	// ex2 익명 구조체 배열
	cars := []struct{ name, color string }{{"520d", "red"}, {"530i", "white"}, {"528i", "blue"}}

	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v) \n", c.name, c.color, c)
	}
}
