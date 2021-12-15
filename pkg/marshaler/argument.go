package marshaler

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/alvarogf97/commandarrgh/pkg/utils"
	"github.com/alvarogf97/commandarrgh/pkg/validations"
)

// It contains every attribute that is
// need to found the argument value.
// It also contains the type of the
// parsed argument, a validator that
// checks the value is right and a
// default value in case the argument
// is not in the given arguments slice
type argumentMarshalerInfo struct {
	// parser tags
	flag     string
	helptext string
	required bool
	isBool   bool

	// value's manipulation
	defaultValue string
	validator    validations.Validator
	kind         reflect.Kind
}

// Gets the given argument value by searching
// it into the given slice attending the
// provided argument information. This method
// will also cast the founded value to
// the required type. If no value is present
// an error will be returned.
//
// Only string, bool or int are casteable any
// other type will result in an error.
func getArgumentValue(arg argumentMarshalerInfo, args []string) (interface{}, error) {
	// search for the value into the given array
	rawValue := ""
	if index, exists := utils.IndexOf(args, arg.flag); exists {
		if arg.isBool {
			rawValue = "true"
		} else {
			rawValue = args[index+1]
		}
	} else if !exists && arg.required {
		return nil, fmt.Errorf("flag %s is required", arg.flag)
	}

	// set the default value if exists and no value
	// had not been founded during the searching
	if rawValue == "" && arg.defaultValue != "" {
		rawValue = arg.defaultValue
	}

	// cast the value to the given one into
	// the argument information
	switch arg.kind {
	case reflect.String:
		return rawValue, nil
	case reflect.Bool:
		return rawValue == "true", nil
	case reflect.Int:
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, fmt.Errorf("invalid value %s for flag %s %s", rawValue, arg.flag, arg.helptext)
		}
		return value, nil
	default:
		return nil, fmt.Errorf("unsuported type `%s`", arg.kind)
	}
}
