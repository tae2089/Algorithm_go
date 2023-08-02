package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		got := isValid("()[]{}")
		assert.Equal(t, true, got)
	})
	t.Run("case2", func(t *testing.T) {
		got := isValid("()")
		assert.Equal(t, true, got)
	})
	t.Run("case1", func(t *testing.T) {
		got := isValid("(]")
		assert.Equal(t, false, got)
	})
}
