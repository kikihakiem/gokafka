package my_helper

import (
	"strconv"
	"strings"
)

const (
	EmptyStr      = ""
	NonIntegerStr = "NN"
)

type StringToIntError struct {
	message   string // description of the error
	errorType string // type of the error
}

func (e *StringToIntError) Error() string {
	return e.message
}

func (e *StringToIntError) Type() string {
	return e.errorType
}

// Extract the number from path
func GetNumber(raw *string) (number int, err *StringToIntError) {
	param := strings.Trim(*raw, "/")
	if param == "" {
		err = &StringToIntError{"String is empty.", EmptyStr}
		return
	}

	number, convertError := strconv.Atoi(param)
	if convertError != nil {
		err = &StringToIntError{"String is non-integer.", NonIntegerStr}
	}

	return
}
