package validations

// Maps each validator name with his
// validator method
var MAPPING = map[string](Validator){
	"path":      pathValidator,
	"stat-path": statPathValidator,
}

// Validator definition
type Validator func(interface{}) (interface{}, error)

// Builds a validator by composing all of
// them into a single function in no
// specified order. That it to say each
// validator must be atomic.
func ComposeValidators(validators []Validator) Validator {
	return func(data interface{}) (interface{}, error) {
		for _, validator := range validators {
			var err error
			if data, err = validator(data); err != nil {
				return nil, err
			}
		}
		return data, nil
	}
}
