package validation

type Validator interface {
	// Validate validates the given value.
	//
	// The performed checks and the supported value types depend on the specific Validator interface implementation.
	// But every implementor must support at least structs, given by value.
	//
	// If all checks are passed, it returns nil.
	// Otherwise, a sole validation Error is returned or a multi-error consisting of one or more validation Error(s).
	Validate(i interface{}) error
}
