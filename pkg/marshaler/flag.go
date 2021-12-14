package marshaler

// Tags are the way the data can be
// marshalled into a given struct.
const (
	TAG_FLAG       = "flag"
	TAG_HELPTEXT   = "helptext"
	TAG_BINDING    = "binding"
	TAG_DEFAULT    = "default"
	TAG_VALIDATORS = "validators"

	TAG_BINDING_REQUIRED = "required"
	TAG_BINDING_IS_BOOL  = "bool"
)
