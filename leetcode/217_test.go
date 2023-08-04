package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		got := containsDuplicate([]int{1, 2, 3, 1})
		assert.Equal(t, true, got)
	})
	t.Run("case2", func(t *testing.T) {
		got := containsDuplicate([]int{1, 2, 3, 4})
		assert.Equal(t, false, got)
	})
	t.Run("case3", func(t *testing.T) {
		got := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
		assert.Equal(t, true, got)
	})
}
