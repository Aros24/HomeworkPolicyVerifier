package policy

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicy_Validate(t *testing.T) {
	tests := []struct {
		name              string
		policyFile        string
		expectedErrorType error
	}{
		{
			name:              "FullyDefinedValidPolicy",
			policyFile:        "full_access_policy.json",
			expectedErrorType: nil,
		},
		{
			name:              "ValidPolicyWithMultipleStatements",
			policyFile:        "complex_policy.json",
			expectedErrorType: nil,
		},
		{
			name:              "PolicyWithUnsupportedVersion",
			policyFile:        "invalid_version_policy.json",
			expectedErrorType: PolicyVersionError(""),
		},
		{
			name:              "PolicyWithInvalidEffect",
			policyFile:        "policy_with_invalid_effect.json",
			expectedErrorType: PolicyEffectValueError(""),
		},
		{
			name:              "PolicyWithoutName",
			policyFile:        "policy_without_name.json",
			expectedErrorType: PolicyNameMissingError(""),
		},
		{
			name:              "PolicyWithInvalidName",
			policyFile:        "policy_with_invalid_name.json",
			expectedErrorType: PolicyNamePaternError(""),
		},
		{
			name:              "PolicyWithoutResource",
			policyFile:        "policy_without_resource.json",
			expectedErrorType: PolicyMissingResourceError(""),
		},
		{
			name:              "PolicyWithoutAction",
			policyFile:        "policy_without_action.json",
			expectedErrorType: PolicyMissingActionError(""),
		},
		{
			name:              "InvalidPolicyNameProperty",
			policyFile:        "invalid_policy_name_property.json",
			expectedErrorType: PolicyNameMissingError(""),
		},
		{
			name:              "InvalidVersionProperty",
			policyFile:        "invalid_version_property.json",
			expectedErrorType: PolicyVersionError(""),
		},
		{
			name:              "MissingRequiredProperty",
			policyFile:        "missing_required_property.json",
			expectedErrorType: PolicyNameMissingError(""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join("../testdata/policyvalidator", test.policyFile)

			data, err := os.ReadFile(path)
			assert.NoError(t, err, "Failed to read %s", test.policyFile)

			var policy Policy
			err = json.Unmarshal(data, &policy)
			assert.NoError(t, err, "Failed to unmarshal policy from %s", test.policyFile)

			err = policy.ValidatePolicy()

			if test.expectedErrorType == nil {
				assert.NoError(t, err, "expected no error")
			} else {
				if assert.Error(t, err, "expected an error") {
					assert.IsType(t, test.expectedErrorType, err, "error type mismatch")
				}
			}
		})
	}
}
func TestPolicyDocument_HasSpecificResources(t *testing.T) {
	tests := []struct {
		name       string
		policyFile string
		wantErr    bool
	}{
		{
			name:       "SpecificResources",
			policyFile: "specific_resources_policy.json",
			wantErr:    true,
		},
		{
			name:       "WildcardResource",
			policyFile: "wildcard_resource_policy.json",
			wantErr:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join("../testdata/policyvalidator", test.policyFile)
			data, err := os.ReadFile(path)
			assert.NoError(t, err, "Failed to read file: %s", test.policyFile)

			var policyDoc PolicyDocument
			err = json.Unmarshal(data, &policyDoc)
			assert.NoError(t, err, "Failed to unmarshal policy document: %s", test.policyFile)

			if test.wantErr {
				assert.True(t, policyDoc.HasSpecificResources(), "Expected specific resources for test: %s", test.name)
			} else {
				assert.False(t, policyDoc.HasSpecificResources(), "Expected wildcard resource for test: %s", test.name)
			}
		})
	}
}
