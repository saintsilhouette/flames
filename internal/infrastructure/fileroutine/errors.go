package fileroutine

import "fmt"

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

var ManagerCreationError = CustomError{Message: "manager was not created!"}
var DirectoryCreationError = CustomError{Message: "directory was not created!"}
var ExistanceUncertaintyError = CustomError{Message: "unable to check directory!"}
var FileCreationError = CustomError{Message: "file was not created!"}
var ImageEncodeError = CustomError{Message: "image was not encoded!"}
var DirectoryReadingError = CustomError{Message: "directory was not red!"}
