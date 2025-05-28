package http

import (
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/common"
)

type Validator interface {
	validate() error
	getFieldName() string
}

func ValidateAll(validators ...Validator) error {
	for _, v := range validators {
		if err := v.validate(); err != nil {
			return fmt.Errorf("bad input (%s): %w", v.getFieldName(), err)
		}
	}
	return nil
}

type StringValidator struct {
	fieldName  string
	value      string
	validators []func(string) error
}

func NewStringValidator(fieldName string, value string, validators ...func(string) error) *StringValidator {
	v := StringValidator{}
	v.fieldName = fieldName
	v.validators = validators
	return &v
}

func (v *StringValidator) getFieldName() string {
	return v.fieldName
}

func (v *StringValidator) validate() error {
	for _, fn := range v.validators {
		if err := fn(v.value); err != nil {
			return err
		}
	}
	return nil
}

// Validator for Ethereum transaction hashes
func IsEthereumTxHash() func(string) error {
	return func(value string) error {
		if !strings.HasPrefix(value, "0x") || len(value) != 66 {
			return fmt.Errorf("invalid transaction hash format")
		}
		return nil
	}
}

// Validator for Ethereum addresses
func IsEthereumAddress() func(string) error {
	return func(value string) error {
		if !common.IsHexAddress(value) {
			return fmt.Errorf("this field must be a valid Ethereum address")
		}
		return nil
	}
}

// Checks if a string is not blank.
func NotBlank() func(string) error {
	return func(value string) error {
		if utf8.RuneCountInString(value) < 1 {
			return fmt.Errorf("this field cannot be blank")
		}
		return nil
	}
}

// Checks if a string has at least n chars.
func CheckMinChars(minChars int) func(string) error {
	return func(value string) error {
		if utf8.RuneCountInString(value) <= minChars {
			return fmt.Errorf("this field must be at least %d characters long", minChars)
		}
		return nil
	}
}

// Checks if a string has at most n chars.
func CheckMaxChars(maxChars int) func(string) error {
	return func(value string) error {
		if utf8.RuneCountInString(value) > maxChars {
			return fmt.Errorf("this field must be at most %d characters long", maxChars)
		}
		return nil
	}
}

// Checks if value is a valid email.
func IsEmail() func(string) error {
	return func(value string) error {
		addr, err := mail.ParseAddress(value)
		if err != nil || addr.Address != value {
			return fmt.Errorf("this field must be a valid email address")
		}
		return nil
	}
}
