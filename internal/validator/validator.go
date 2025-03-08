package validator

import (
	"mime/multipart"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	FieldErrors map[string][]string
	FormErrors  []string
}

func (v *Validator) IsValid() bool {
	return len(v.FieldErrors) == 0 && len(v.FormErrors) == 0
}

func (v *Validator) AddFormError(message string) {
	if v.FormErrors == nil {
		v.FormErrors = make([]string, 0)
	}

	v.FormErrors = append(v.FormErrors, message)
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string][]string)
	}

	_, exists := v.FieldErrors[key]

	if exists {
		v.FieldErrors[key] = append(v.FieldErrors[key], message)
	} else {
		v.FieldErrors[key] = []string{message}
	}
}

func (v *Validator) CheckField(isOk bool, key, message string) {
	if !isOk {
		v.AddFieldError(key, message)
	}
}

func FileNotBlank(fileHeader multipart.FileHeader) bool {
	return fileHeader.Size != 0
}

func FileTypeNotAllowed(fileHeader multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	return allowedTypes[fileHeader.Header.Get("Content-Type")]
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func SameValue(valueOne string, valueTwo string) bool {
	return strings.EqualFold(valueOne, valueTwo)
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
