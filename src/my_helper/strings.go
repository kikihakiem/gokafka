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
func GetNumber(raw *string) (int, *StringToIntError) {
	param := strings.Trim(*raw, "/")
	if param == "" {
		return 0, &StringToIntError{"String is empty.", EmptyStr}
	}

	number, err := strconv.Atoi(param)
	if err != nil {
		return 0, &StringToIntError{"String is non-integer.", NonIntegerStr}
	}

	return number, nil
}
