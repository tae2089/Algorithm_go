package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumWealth(t *testing.T) {
	t.Run("case1 ", func(t *testing.T) {
		got := maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}})
		assert.Equal(t, got, 6)
	})

	t.Run("case2", func(t *testing.T) {
		got := maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}})
		assert.Equal(t, got, 10)
	})

	t.Run("case3", func(t *testing.T) {
		got := maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})
		assert.Equal(t, got, 17)
	})
}
