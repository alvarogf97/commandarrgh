package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitCommand(t *testing.T) {
	assert := require.New(t)

	t.Run("test_split_success", func(t *testing.T) {
		command := "commit"
		args := "-m 'first commit'"

		rcommand, rargs := SplitCommand(fmt.Sprintf("%s %s", command, args))

		assert.Equal(command, rcommand)
		assert.Equal(args, rargs)
	})

}
