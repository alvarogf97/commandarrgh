package validations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComposeValidators(t *testing.T) {
	assert := require.New(t)

	t.Run("test_compose_validator_that_succeeds", func(t *testing.T) {
		succeedValidator := func(v interface{}) (interface{}, error) { return v, nil }
		composedValidator := ComposeValidators([]Validator{succeedValidator})

		_, err := composedValidator("")

		assert.NoError(err)
	})

	t.Run("test_compose_validator_that_fails", func(t *testing.T) {
		ferr := fmt.Errorf("Error")
		succeedValidator := func(v interface{}) (interface{}, error) { return v, nil }
		faliedValidator := func(v interface{}) (interface{}, error) { return v, ferr }
		composedValidator := ComposeValidators([]Validator{succeedValidator, faliedValidator})

		_, err := composedValidator("")

		assert.Error(err, ferr)
	})
}
