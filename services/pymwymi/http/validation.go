package http

import (
	"fmt"
	"net/mail"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

type Validator interface {
	validate() error
	getFieldName() string
}

func ValidateAll(validators ...Validator) *pymwymi.Error {
	for _, v := range validators {
		if err := v.validate(); err != nil {
			return pymwymi.Errorf(pymwymi.ErrBadInput, "bad input (%s): %s", v.getFieldName(), err.Error())
		}
	}
	return nil
}

type IntegerValidator struct {
	fieldName  string
	value      string
	validators []func(string) error
}

type StringValidator struct {
	fieldName  string
	value      string
	validators []func(string) error
}

// ---------------------------------
// Integer validators
func NewIntegerValidator(fieldName string, value string, validators ...func(string) error) *IntegerValidator {
	v := IntegerValidator{}
	v.fieldName = fieldName
	v.validators = validators
	return &v
}

func (v *IntegerValidator) getFieldName() string {
	return v.fieldName
}

func (v *IntegerValidator) validate() error {
	for _, fn := range v.validators {
		if err := fn(v.value); err != nil {
			return err
		}
	}
	return nil
}

func CheckIsInt64() func(string) error {
	return func(value string) error {
		if _, err := strconv.ParseInt(value, 10, 64); err != nil {
			return fmt.Errorf("this field must be an integer")
		}
		return nil
	}
}

// can take in a string or an int64
func CheckIsBetween[T string | int64](min, max int64) func(T) error {
	return func(value T) error {
		var intValue int64
		var err error
		switch v := any(value).(type) {
		case int64:
			intValue = v
		case string:
			intValue, err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				return fmt.Errorf("this field must be an integer")
			}
		default:
			return fmt.Errorf("unsupported type")
		}
		if intValue < min {
			return fmt.Errorf("this field must be greater than %d", min)
		}
		if intValue > max {
			return fmt.Errorf("this field must be less than %d", max)
		}
		return nil
	}
}

// ---------------------------------
// String validators
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
		if utf8.RuneCountInString(value) < minChars {
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

// Checks if value is a valid email.

func IsInList() func(string, ...string) error {
	return func(valueToCheck string, values ...string) error {
		if !slices.Contains(values, valueToCheck) {
			return fmt.Errorf("value not included in list of possible values")
		}
		return nil
	}
}
