package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a, b int
	fmt.Fscanln(reader, &a, &b) // 한줄 입력, 띄어쓰기로 구분

	fmt.Fprintln(writer, a, b)
	fmt.Fprintln(writer, a+b)
	// fmt.Println("result : ", a, b)
}
