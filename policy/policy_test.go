package policy

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestPolicy_Validate(t *testing.T) {
	tests := []struct {
		name              string
		policyFile        string
		expectError       bool
		expectedErrorType reflect.Type
	}{
		{
			name:              "FullyDefinedValidPolicy",
			policyFile:        "full_access_policy.json",
			expectError:       false,
			expectedErrorType: nil,
		},
		{
			name:              "ValidPolicyWithMultipleStatements",
			policyFile:        "complex_policy.json",
			expectError:       false,
			expectedErrorType: nil,
		},
		{
			name:              "PolicyWithUnsupportedVersion",
			policyFile:        "invalid_version_policy.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyVersionError("")),
		},
		{
			name:              "PolicyWithInvalidEffect",
			policyFile:        "policy_with_invalid_effect.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyEffectValueError("")),
		},
		{
			name:              "PolicyWithoutName",
			policyFile:        "policy_without_name.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyNameMissingError("")),
		},
		{
			name:              "PolicyWithInvalidName",
			policyFile:        "policy_with_invalid_name.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyNamePaternError("")),
		},
		{
			name:              "PolicyWithoutResource",
			policyFile:        "policy_without_resource.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyMissingResourceError("")),
		},
		{
			name:              "PolicyWithoutAction",
			policyFile:        "policy_without_action.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyMissingActionError("")),
		},
		{
			name:              "InvalidPolicyNameProperty",
			policyFile:        "invalid_policy_name_property.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyNameMissingError("")),
		},
		{
			name:              "InvalidVersionProperty",
			policyFile:        "invalid_version_property.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyVersionError("")),
		},
		{
			name:              "MissingRequiredProperty",
			policyFile:        "missing_required_property.json",
			expectError:       true,
			expectedErrorType: reflect.TypeOf(PolicyNameMissingError("")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join("../testdata/policyvalidator", test.policyFile)

			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", test.policyFile, err)
			}

			var policy Policy
			err = json.Unmarshal(data, &policy)
			if err != nil {
				t.Fatalf("Failed to unmarshal policy from %s: %v", test.policyFile, err)
			}

			err = policy.ValidatePolicy()

			if test.expectError {
				if err == nil {
					t.Errorf("Policy.Validate() expected error, got none")
				} else if reflect.TypeOf(err) != test.expectedErrorType {
					t.Errorf("Policy.Validate() error = %T, want %T", err, test.expectedErrorType)
				}
			} else if err != nil {
				t.Errorf("Policy.Validate() unexpected error = %v", err)
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
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", test.policyFile, err)
			}

			var policyDoc PolicyDocument
			if err := json.Unmarshal(data, &policyDoc); err != nil {
				t.Fatalf("Failed to unmarshal policy document from %s: %v", test.policyFile, err)
			}

			if got := policyDoc.HasSpecificResources(); got != test.wantErr {
				t.Errorf("PolicyDocument.HasSpecificResources() = %v, want %v", got, test.wantErr)
			}
		})
	}
}
