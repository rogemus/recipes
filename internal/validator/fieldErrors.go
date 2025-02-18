package validator

import "fmt"

type fieldErrors struct{}

func (f fieldErrors) ErrNotBlank() string {
	return "Field cannot be empty"
}

func (f fieldErrors) ErrNotEmail() string {
	return "Value is not valid email adress"
}

func (f fieldErrors) ErrMaxLength(n int) string {
	return fmt.Sprintf("Value must have less then %d characters", n)
}

func (f fieldErrors) ErrMinLength(n int) string {
	return fmt.Sprintf("Value must have at last %d characters", n)
}

func (f fieldErrors) ErrPassNotSame() string {
	return "Password does not match"
}

type formErros struct{}

func (f formErros) ErrInvalidCredentials() string {
	return "Invalid credentials"
}

var FormErros = formErros{}
var FieldErr = fieldErrors{}
