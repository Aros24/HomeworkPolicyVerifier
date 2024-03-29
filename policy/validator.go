package policy

import (
	"regexp"
)

func (policy *Policy) ValidatePolicy() error {
	if policy.PolicyName == "" {
		return PolicyNameMissingError(ErrPolicyNameRequired)
	}
	matched, _ := regexp.MatchString(`^[\w+=,.@-]{1,128}$`, policy.PolicyName)
	if !matched {
		return PolicyNamePaternError(ErrPolicyNameInvalidPattern)
	}

	return policy.PolicyDocument.validateDocument()
}

func (policyDocument *PolicyDocument) validateDocument() error {
	if policyDocument.Version != "2012-10-17" && policyDocument.Version != "2008-10-17" {
		return PolicyVersionError(ErrPolicyVersionUnsupported)
	}

	for _, stmt := range policyDocument.Statement {
		if err := stmt.validateStatement(); err != nil {
			return err
		}
	}
	return nil
}

func (statement *Statement) validateStatement() error {
	if statement.Effect != "Allow" && statement.Effect != "Deny" {
		return PolicyEffectValueError(ErrStatementEffectInvalid)
	}

	if len(statement.Action) == 0 {
		return PolicyMissingActionError(ErrStatementActionMissing)
	}

	if len(statement.Resource) == 0 {
		return PolicyMissingResourceError(ErrStatementResourceMissing)
	}
	return nil
}

func (policyDocument *PolicyDocument) HasSpecificResources() bool {
	for _, statement := range policyDocument.Statement {
		if len(statement.Resource) == 1 && statement.Resource[0] == "*" {
			return false
		}
	}
	return true
}
