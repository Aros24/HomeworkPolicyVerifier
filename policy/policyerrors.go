package policy

const (
	ErrPolicyNameRequired          = "policy name is required"
	ErrPolicyNameInvalidPattern    = "policy name does not match the required pattern"
	ErrPolicyVersionUnsupported    = "unsupported policy version"
	ErrStatementStructureInvalid   = "invalid structure of statement"
	ErrStatementEffectInvalid      = "invalid effect: must be 'Allow' or 'Deny'"
	ErrStatementActionMissing      = "at least one action is required"
	ErrStatementResourceMissing    = "at least one resource is required"
	ErrSpecificResourceRequirement = "specific resource requirement not met"
	errUnmarshalSlicedString       = "invalid format for sliced string"
)

type PolicyNamePaternError string

func (e PolicyNamePaternError) Error() string {
	return string(e)
}

type PolicyNameMissingError string

func (e PolicyNameMissingError) Error() string {
	return string(e)
}

type PolicyVersionError string

func (e PolicyVersionError) Error() string {
	return string(e)
}

type PolicyEffectValueError string

func (e PolicyEffectValueError) Error() string {
	return string(e)
}

type PolicyMissingActionError string

func (e PolicyMissingActionError) Error() string {
	return string(e)
}

type PolicyMissingResourceError string

func (e PolicyMissingResourceError) Error() string {
	return string(e)
}

type PolicyStatementInvalidError string

func (e PolicyStatementInvalidError) Error() string {
	return string(e)
}

type UnmarshalError string

func (e UnmarshalError) Error() string {
	return string(e)
}
