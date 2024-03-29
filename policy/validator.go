package policy

import (
	"errors"
	"regexp"
	"strconv"
)

func (policy *Policy) ValidatePolicy() error {
	if policy.PolicyName == "" {
		return errors.New("policy name is required")
	}
	matched, _ := regexp.MatchString(`^[\w+=,.@-]{1,128}$`, policy.PolicyName)
	if !matched {
		return errors.New("policy name does not match the required pattern")
	}

	return policy.PolicyDocument.validateDocument()
}

func (policyDocument *PolicyDocument) validateDocument() error {
	if policyDocument.Version != "2012-10-17" && policyDocument.Version != "2008-10-17" {
		return errors.New("unsupported policy version: " + policyDocument.Version)
	}

	for i, stmt := range policyDocument.Statement {
		if err := stmt.validateStatement(); err != nil {
			return errors.New("statement " + strconv.Itoa(i) + ": " + err.Error())
		}
	}
	return nil
}

func (statement *Statement) validateStatement() error {
	if statement.Effect != "Allow" && statement.Effect != "Deny" {
		return errors.New("invalid effect: " + statement.Effect)
	}

	if len(statement.Action) == 0 {
		return errors.New("at least one action is required")
	}

	if len(statement.Resource) == 0 {
		return errors.New("at least one resource is required")
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
