package utils

import (
	"io/fs"
	"os"
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	tempFile, err := os.Create("example.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	content := `{"test": "data"}`
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	tests := []struct {
		name            string
		input           string
		want            string
		expectError     bool
		expectedErrType reflect.Type
	}{
		{
			name:            "ValidFile",
			input:           tempFile.Name(),
			want:            content,
			expectError:     false,
			expectedErrType: nil,
		},
		{
			name:            "InvalidFile",
			input:           "nonexistent.json",
			want:            "",
			expectError:     true,
			expectedErrType: reflect.TypeOf((*fs.PathError)(nil)),
		},
		{
			name:            "DirectString",
			input:           `{"direct": "input"}`,
			want:            `{"direct": "input"}`,
			expectError:     false,
			expectedErrType: nil,
		},
		{
			name:            "MalformedJsonInput",
			input:           `{"malformedJson": true,}`,
			want:            "",
			expectError:     true,
			expectedErrType: reflect.TypeOf(InputWrongJsonFormatError("")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ReadInput(test.input)

			if test.expectError {
				if err == nil {
					t.Errorf("ReadInput() expected error, got none")
				} else if reflect.TypeOf(err) != test.expectedErrType {
					t.Errorf("ReadInput() error = %v, want %v", reflect.TypeOf(err), test.expectedErrType)
				}
			} else if err != nil {
				t.Errorf("ReadInput() unexpected error = %v", err)
			}
			if string(got) != test.want {
				t.Errorf("ReadInput() = %v, want %v", string(got), test.want)
			}
		})
	}
}
