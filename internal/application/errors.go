package application

import "fmt"

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

var InvalidHeightError = CustomError{Message: "invalid height value!"}
var NegativeHeightError = CustomError{Message: "negative height value!"}
var InvalidWidthError = CustomError{Message: "invalid width value!"}
var NegativeWidthError = CustomError{Message: "negative width value!"}
var InvalidIterationsError = CustomError{Message: "invalid iterations value!"}
var NegativeIterationsError = CustomError{Message: "negative iterations value!"}
var InvalidGoroutinesError = CustomError{Message: "invalid iterations value!"}
var NegativeGoroutinesError = CustomError{Message: "negative goroutines value!"}
