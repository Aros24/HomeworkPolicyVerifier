package utils

const (
	ErrWrongJsonFormat    = "provided input is not in valid JSON format"
	ErrWrongFileExtension = "provided file with wrong extension"
)

type InputWrongJsonFormatError string

func (e InputWrongJsonFormatError) Error() string {
	return string(e)
}

type WrongFileExtension string

func (e WrongFileExtension) Error() string {
	return string(e)
}
