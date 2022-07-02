package lv1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNotExistNum(t *testing.T) {

	t.Run("테스트1", func(t *testing.T) {
		//given
		numbers := []int{1, 2, 3, 4, 6, 7, 8, 0}

		//when
		rslt := AddNotExistNum(numbers)

		//then: 14
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.Equal(t, rslt, 14)
	})

	t.Run("테스트2", func(t *testing.T) {
		//given
		numbers := []int{5, 8, 4, 0, 6, 7, 9}

		//when
		rslt := AddNotExistNum(numbers)

		//then: 14
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.Equal(t, rslt, 6)
	})
}
