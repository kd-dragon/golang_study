package lv1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberKtimes(t *testing.T) {

	t.Run("테스트1", func(t *testing.T) {
		//given
		array := []int{1, 5, 2, 6, 3, 7, 4}
		commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}

		//when
		rslt := NumberKtimes(array, commands)

		//then: [5, 6, 3]
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.ElementsMatch(t, rslt, []int{5, 6, 3})
	})
}
