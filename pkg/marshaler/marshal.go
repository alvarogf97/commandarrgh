package marshaler

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alvarogf97/commandarrgh/pkg/utils"
	"github.com/alvarogf97/commandarrgh/pkg/validations"
	"github.com/google/shlex"
)

// Marshall a given string slice into
// the given struct. This method will
// iterate over the given slice and
// check every `flag` tags in order
// to fill it with the correct value.
//
// Example:
//
// type myArgs struct {
//     Directory  string `flag:"-d" helptext:"directory" binding:"false" default:"./" validators:"path"`
// 	   ShowHidden bool   `flag:"-h" helptext:"shows hidden folders" binding:"isBool"`
// }
//
// argsList := ["-d", "/memes", "-h"]
// args := &myArgs{}
//
// MarshallCommandArguments(argsList, args)
//
// args.Directory must contains "/memes"
// args.ShowHidden must be true
func MarshalCommandArguments(args string, v interface{}) error {
	// splits the given args into an array
	argsArray, err := shlex.Split(args)
	if err != nil {
		return err
	}

	// checks v is a struct pointer
	ptr := reflect.ValueOf(v)
	if ptr.Kind() != reflect.Ptr {
		return fmt.Errorf("%s is not a pointer to a struct", v)
	}

	// iterates over the struct elements in order
	// to identify each field and fill every
	// `argumentMarshallerInfo` in order to
	// identify its value
	rvalue := ptr.Elem()
	rtype := rvalue.Type()
	for i := 0; i < rtype.NumField(); i++ {
		arg := argumentMarshalerInfo{kind: rvalue.Field(i).Kind()}

		if tagValue, ok := rtype.Field(i).Tag.Lookup(TAG_FLAG); ok {
			arg.flag = tagValue
		} else {
			// TAG_FLAG is mandatory, if it is not present
			// an error will be returned
			return fmt.Errorf(
				"field %s has no flag value", rtype.Field(i).Name)
		}

		// set argument helptext
		if tagValue, ok := rtype.Field(i).Tag.Lookup(TAG_HELPTEXT); ok {
			arg.helptext = tagValue
		}

		// set argument default value. It must be casteable
		// to the right struct declared type
		if tagValue, ok := rtype.Field(i).Tag.Lookup(TAG_DEFAULT); ok {
			arg.defaultValue = tagValue
		}

		// set validator function by composing all of them
		// in no mandatory order. Those ones will be executed
		// before setting the value. If no validator is present,
		// an empty validator that returns the value will be
		// assigned as default
		if tagValue, ok := rtype.Field(i).Tag.Lookup(TAG_VALIDATORS); ok {
			validatorNames := strings.Split(tagValue, ",")
			validators := make([]validations.Validator, len(validatorNames))
			for i, validatorName := range validatorNames {
				// checks the specified validator is registered
				// into the `VALIDATOR_MAPPING`
				if validator, exists := validations.MAPPING[validatorName]; !exists {
					return fmt.Errorf("validator: %s does not exists", tagValue)
				} else {
					validators[i] = validator
				}
			}

			// compose validators into a single one
			arg.validator = validations.ComposeValidators(validators)
		} else {
			// set empty validator that returns the value
			arg.validator = func(value interface{}) (interface{}, error) { return value, nil }
		}

		// check binding options in order to set the final
		// value or the given default
		if tagValue, ok := rtype.Field(i).Tag.Lookup(TAG_BINDING); ok {
			splittedValue := strings.Split(tagValue, ",")
			_, arg.required = utils.IndexOf(splittedValue, TAG_BINDING_REQUIRED)
			_, arg.isBool = utils.IndexOf(splittedValue, TAG_BINDING_IS_BOOL)
		}

		// get argument value into his own type
		fieldValue, err := getArgumentValue(arg, argsArray)
		if err != nil {
			return err
		}

		// now the value is in his own type, the
		// validator will be executed
		validatedFieldValue, err := arg.validator(fieldValue)
		if err != nil {
			return err
		}

		// finally fill the value into his field
		rvalue.Field(i).Set(reflect.ValueOf(validatedFieldValue))
	}
	return nil
}
