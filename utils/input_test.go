package utils

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
		name              string
		input             string
		want              string
		expectedErrorType error
	}{
		{
			name:              "ValidFile",
			input:             tempFile.Name(),
			want:              content,
			expectedErrorType: nil,
		},
		{
			name:              "InvalidFileContent",
			input:             "../testdata/readinput/invalidfilecontent.json",
			want:              "",
			expectedErrorType: InputWrongJsonFormatError(""),
		},
		{
			name:              "InvalidFile",
			input:             "nonexistent.json",
			want:              "",
			expectedErrorType: (*fs.PathError)(nil),
		},
		{
			name:              "WrongFileFormat",
			input:             "../testdata/readinput/blanknotepad.txt",
			want:              "",
			expectedErrorType: WrongFileExtension(""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ReadInput(test.input)

			if test.expectedErrorType == nil {
				assert.NoError(t, err, "expected no error")
			} else {
				if assert.Error(t, err, "expected an error") {
					assert.IsType(t, test.expectedErrorType, err, "error type mismatch")
				}
			}

			if test.expectedErrorType != InputWrongJsonFormatError("") {
				assert.Equal(t, test.want, string(got))
			}
		})
	}
}
