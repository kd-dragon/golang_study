package lv1

import "sort"

// https://programmers.co.kr/learn/courses/30/lessons/42748?language=go

func NumberKtimes(array []int, commands [][]int) []int {

	var answers []int
	var tarry []int

	for _, cmd := range commands {
		tarry = make([]int, len(array))
		copy(tarry, array)
		tarry = tarry[cmd[0]-1 : cmd[1]]
		sort.Ints(tarry)
		answers = append(answers, tarry[cmd[2]-1])
	}

	return answers
}
