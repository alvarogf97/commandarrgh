package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayIndexOf(t *testing.T) {
	assert := require.New(t)

	t.Run("test_found", func(t *testing.T) {
		value := "bonk"
		array := []string{"gorillaz", "bonk", "dog"}
		position, found := IndexOf(array, value)

		assert.True(found)
		assert.Equal(1, position)
	})

	t.Run("test_not_found", func(t *testing.T) {
		value := "cat"
		array := []string{"gorillaz", "bonk", "dog"}
		position, found := IndexOf(array, value)

		assert.False(found)
		assert.Equal(-1, position)
	})
}
