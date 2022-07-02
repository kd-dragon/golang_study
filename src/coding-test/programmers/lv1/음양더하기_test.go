package lv1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumAbsolutes(t *testing.T) {

	t.Run("테스트1", func(t *testing.T) {
		//given
		absolutes := []int{4, 7, 12}
		signs := []bool{true, false, true}

		//when
		rslt := SumAbsolutes(absolutes, signs)

		//then: 9
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.Equal(t, rslt, 9)
	})

	t.Run("테스트2", func(t *testing.T) {
		//given
		absolutes := []int{1, 2, 3}
		signs := []bool{false, false, true}

		//when
		rslt := SumAbsolutes(absolutes, signs)

		//then: 14
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.Equal(t, rslt, 0)
	})
}
