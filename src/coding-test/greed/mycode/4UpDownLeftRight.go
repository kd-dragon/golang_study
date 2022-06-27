// 상하좌우
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handler(cmd string, position []int, limit int) {
	switch cmd {
	case "R":
		if position[1] < limit {
			position[1]++
		}
	case "D":
		if position[0] < limit {
			position[0]++
		}
	case "L":
		if position[1] > 1 {
			position[1]--
		}
	case "U":
		if position[0] > 1 {
			position[0]--
		}
	}
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	position := []int{1, 1}

	var N int

	fmt.Scanln(&N)

	scn.Scan()

	plans := strings.Split(scn.Text(), " ")

	for _, v := range plans {
		handler(v, position, N)
	}

	fmt.Println(position)
}
