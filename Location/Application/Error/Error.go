package Error

import "fmt"

type DataValidationError struct {
	inner error
}

func NewDataValidationError(inner error) *DataValidationError {
	return &DataValidationError{inner: inner}
}

func (v DataValidationError) Error() string {
	return fmt.Sprintf("Data was invalid, reason: %s", v.inner.Error())
}
