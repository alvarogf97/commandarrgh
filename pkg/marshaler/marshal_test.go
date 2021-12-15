package marshaler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalComandArguments(t *testing.T) {
	assert := require.New(t)

	type mockArgs struct {
		Dst            string `flag:"-dst" helptext:"file's directory" binding:"required" validators:"stat-path"`
		Src            string `flag:"-src" helptext:"file's directory" binding:"required" validators:"stat-path"`
		AlsoGetHiddens bool   `flag:"-h" helptext:"include hidden folder's files" binding:"bool" default:"false" `
	}

	type mockArgsNoTag struct {
		mockArgs
		Untagged string
	}

	type mockArgsInvalidValidator struct {
		Invalid string `flag:"-i" validators:"fake"`
	}

	type mockArgsInvalidValue struct {
		Invalid int `flag:"-"`
	}

	t.Run("test_marshal_command_success", func(t *testing.T) {
		data := "-dst ./ -src ./ -h"
		args := &mockArgs{}

		err := MarshalCommandArguments(data, args)

		assert.NoError(err)
		assert.True(args.AlsoGetHiddens)
	})

	t.Run("test_marshal_command_fail_not_a_struct_pointer", func(t *testing.T) {
		err := MarshalCommandArguments("", 3)
		assert.Error(err)
	})

	t.Run("test_marshal_command_fail_no_tag_flag", func(t *testing.T) {
		data := "-dst ./ -src ./ -h"
		args := &mockArgsNoTag{}

		err := MarshalCommandArguments(data, args)

		assert.Error(err)
	})

	t.Run("test_marshal_command_fail_validator_not_exists", func(t *testing.T) {
		data := "-i data"
		args := &mockArgsInvalidValidator{}

		err := MarshalCommandArguments(data, args)

		assert.Error(err)
	})

	t.Run("test_marshal_command_fail_getting_argument_value", func(t *testing.T) {
		data := "-i data"
		args := &mockArgsInvalidValue{}

		err := MarshalCommandArguments(data, args)

		assert.Error(err)
	})

	t.Run("test_marshal_command_fail_validator_check", func(t *testing.T) {
		data := "-dst ./ -src ./fake -h"
		args := &mockArgs{}

		err := MarshalCommandArguments(data, args)

		assert.Error(err)
	})

}
