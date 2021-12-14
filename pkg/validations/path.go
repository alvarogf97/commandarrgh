package validations

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

// Checks the given data is a wellformed path
// and normalize it if necessary by converting it
// to an absolute path and replacing backslashes
// for full system compatibility
func pathValidator(data interface{}) (interface{}, error) {

	// checks data is a string
	if reflect.ValueOf(data).Kind() != reflect.String {
		return nil, fmt.Errorf("path validator can only be used on string fields")
	}

	// converts to an absolute path
	// and normalize it
	path := (data).(string)
	asbPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("cannot convert path  to asbpath %s", err.Error())
	}

	return asbPath, nil
}

// Checks the given data is a wellformed path
// and checks that this path exists in the system
func statPathValidator(data interface{}) (interface{}, error) {
	value, err := pathValidator(data)
	if err != nil {
		return nil, err
	}
	path := (value).(string)

	// checks the path exists
	_, err = os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("path %s does not exists", path)
	}

	return path, nil
}
