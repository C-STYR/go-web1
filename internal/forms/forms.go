package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form defines a structure for containing url.Values and associated errors
type Form struct {
	url.Values
	Errors errors
}

// Determines validation of a given form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a Form struct with passed url.Values
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required validates that user has input required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if the field specified has a value (doesn't add an error)
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	// if no value for field, return false
	return x != ""
}

// MinLength checks user input values for minimum length
func (f *Form) MinLength(field string, minLen int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < minLen {
		f.Errors.Add(field, fmt.Sprintf("The field length must be at least %d characters", minLen))
		return false
	}
	return true
}
