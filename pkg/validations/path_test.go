package validations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathValidator(t *testing.T) {
	assert := require.New(t)

	t.Run("test_path_success", func(t *testing.T) {
		path := "./"
		_, err := pathValidator(path)

		assert.NoError(err)
	})

	t.Run("test_path_fail_not_string", func(t *testing.T) {
		path := 10
		_, err := pathValidator(path)

		assert.Error(err)
	})
}

func TestStatPathValidator(t *testing.T) {
	assert := require.New(t)

	t.Run("test_path_success", func(t *testing.T) {
		path := "./"
		_, err := statPathValidator(path)

		assert.NoError(err)
	})

	t.Run("test_path_fail_path_validator", func(t *testing.T) {
		path := 10
		_, err := statPathValidator(path)

		assert.Error(err)
	})

	t.Run("test_path_fail_not_exists", func(t *testing.T) {
		path := "fakepath"
		_, err := statPathValidator(path)

		assert.Error(err)
	})
}
