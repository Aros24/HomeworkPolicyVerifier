package utils

const (
	ErrWrongJsonFormat = "provided input is not in valid JSON format"
)

type InputWrongJsonFormatError string

func (e InputWrongJsonFormatError) Error() string {
	return string(e)
}
