package main

import (
	"fmt"
	"sort"
)

func main() {

	var n, m, k int
	fmt.Scanf("%d %d %d\n", &n, &m, &k)

	var slice []int = make([]int, n)

	for i := 0; i < n; {
		fmt.Scanf("%d", &slice[i])
		i++
	}
	sort.Ints(slice)
	fmt.Println("max, sub :", slice[n-1], slice[n-2])

	var output int
	// Legacy code
	for i := 0; i < m; {
		if k < m {
			output += (slice[n-1] * k)
			output += slice[n-2]
			i += (k + 1)
		} else if k > m {
			output += (slice[n-1] * m)
			i += m
		} else {
			output += (slice[n-1] * k)
			i += k
		}
	}
	fmt.Println("final value ::: ", output)
}
