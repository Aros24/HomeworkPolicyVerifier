package policy

import (
	"encoding/json"
	"regexp"
)

type Policy struct {
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
	Version    string      `json:"Version"`
	Statements []Statement `json:"Statement"`
}

type Statement struct {
	Sid       string      `json:"Sid,omitempty"`
	Effect    string      `json:"Effect"`
	Action    SliceString `json:"Action"`
	Resource  SliceString `json:"Resource"`
	Condition *Condition  `json:"Condition,omitempty"`
}

type SliceString []string

type Condition struct {
	Bool map[string]string `json:"Bool"`
}

const policyNamePattern = `^[\w+=,.@-]{1,128}$`

// Overload if has only one or multiple values (Array)
func (textToMarshall *SliceString) UnmarshalJSON(data []byte) error {
	var singleValue string
	if err := json.Unmarshal(data, &singleValue); err == nil {
		*textToMarshall = []string{singleValue}
		return nil
	}

	var multiValue []string
	if err := json.Unmarshal(data, &multiValue); err == nil {
		*textToMarshall = multiValue
		return nil
	}

	return UnmarshalError(errUnmarshalSlicedString)
}

func UnmarshalPolicy(data []byte) (*Policy, error) {
	var p Policy
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (policy *Policy) ValidatePolicy() error {
	if policy.PolicyName == "" {
		return PolicyNameMissingError(ErrPolicyNameRequired)
	}
	matched, _ := regexp.MatchString(policyNamePattern, policy.PolicyName)
	if !matched {
		return PolicyNamePaternError(ErrPolicyNameInvalidPattern)
	}

	return policy.PolicyDocument.validateDocument()
}

type versionSupport struct {
	supportedVersions []string
}

var supportedVersionInstance = versionSupport{
	supportedVersions: []string{"2012-10-17", "2008-10-17"},
}

func (policy *PolicyDocument) isVersionSupported() bool {
	for _, supportedVersion := range supportedVersionInstance.supportedVersions {
		if policy.Version == supportedVersion {
			return true
		}
	}
	return false
}

func (policyDocument *PolicyDocument) validateDocument() error {
	if !policyDocument.isVersionSupported() {
		return PolicyVersionError(ErrPolicyVersionUnsupported)
	}

	for _, stmt := range policyDocument.Statements {
		if err := stmt.validateStatement(); err != nil {
			return err
		}
	}
	return nil
}

type effectSupport struct {
	supportedEffects []string
}

var supportedEffectInstance = effectSupport{
	supportedEffects: []string{"Allow", "Deny"},
}

func (statement Statement) isEffectSupported() bool {
	for _, supportedEffect := range supportedEffectInstance.supportedEffects {
		if statement.Effect == supportedEffect {
			return true
		}
	}
	return false
}

func (statement *Statement) validateStatement() error {
	if !statement.isEffectSupported() {
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
	for _, statement := range policyDocument.Statements {
		if len(statement.Resource) == 1 && statement.Resource[0] == "*" {
			return false
		}
	}
	return true
}
