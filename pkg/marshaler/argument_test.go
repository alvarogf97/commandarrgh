package marshaler

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetArgumentValue(t *testing.T) {
	assert := require.New(t)

	t.Run("test_get_argument_value_success_string", func(t *testing.T) {
		flag := "-f"
		value := "dog"
		arg := argumentMarshalerInfo{flag: flag, kind: reflect.String}
		args := []string{flag, value}

		result, err := getArgumentValue(arg, args)

		assert.NoError(err)
		assert.Equal(value, result)
	})

	t.Run("test_get_argument_value_success_bool", func(t *testing.T) {
		flag := "-f"
		arg := argumentMarshalerInfo{flag: flag, kind: reflect.Bool, isBool: true}
		args := []string{flag}

		result, err := getArgumentValue(arg, args)

		assert.NoError(err)
		assert.True((result).(bool))
	})

	t.Run("test_get_argument_value_success_int", func(t *testing.T) {
		flag := "-f"
		value := 1
		valuestr := "1"
		arg := argumentMarshalerInfo{flag: flag, kind: reflect.Int}
		args := []string{flag, valuestr}

		result, err := getArgumentValue(arg, args)

		assert.NoError(err)
		assert.Equal(value, result)
	})

	t.Run("test_get_argument_value_success_default_value", func(t *testing.T) {
		flag := "-f"
		value := "dog"
		arg := argumentMarshalerInfo{flag: flag, kind: reflect.String, defaultValue: value}
		args := []string{}

		result, err := getArgumentValue(arg, args)

		assert.NoError(err)
		assert.Equal(value, result)
	})

	t.Run("test_get_argument_value_fail_no_required_value", func(t *testing.T) {
		flag := "-f"
		arg := argumentMarshalerInfo{flag: flag, kind: reflect.String, required: true}
		args := []string{}

		_, err := getArgumentValue(arg, args)

		assert.Error(err)
	})

	t.Run("test_get_argument_value_fail_cast_int", func(t *testing.T) {
		flag := "-f"
		valuestr := "dog"
		arg := argumentMarshalerInfo{flag: flag, kind: reflect.Int}
		args := []string{flag, valuestr}

		_, err := getArgumentValue(arg, args)
		assert.Error(err)
	})

	t.Run("test_get_argument_value_fail_unsuported_type", func(t *testing.T) {
		flag := "-f"
		value := "dog"
		arg := argumentMarshalerInfo{flag: flag, kind: reflect.Array}
		args := []string{flag, value}

		_, err := getArgumentValue(arg, args)

		assert.Error(err)
	})
}
