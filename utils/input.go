package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadInput(inputPathOrString string) ([]byte, error) {
	var data []byte
	var err error

	if strings.HasSuffix(inputPathOrString, ".json") {
		data, err = os.ReadFile(filepath.Clean(inputPathOrString))
		if err != nil {
			return nil, err
		}

		if !isValidJSON(data) {
			return nil, InputWrongJsonFormatError(ErrWrongJsonFormat)
		}
	} else {
		return nil, WrongFileExtension(ErrWrongFileExtension)
	}

	return data, nil
}

func isValidJSON(data []byte) bool {
	var js json.RawMessage
	err := json.Unmarshal(data, &js)
	if err != nil {
		log.Println("JSON validation error:", err)
		return false
	}
	return true
}
