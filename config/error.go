package config

import "fmt"

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

var WidthValueOverflow = CustomError{Message: "width value overflow!"}
var HeightValueOverflow = CustomError{Message: "height value overflow!"}
var SymmetryValueOverflow = CustomError{Message: "symmetry value overflow!"}
