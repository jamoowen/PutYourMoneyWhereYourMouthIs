package pymwymi

import (
	"errors"
	"fmt"
)

type ErrorCode string

const (
	ErrUserNotFound      ErrorCode = "user_not_found"
	ErrUserAlreadyExists ErrorCode = "user_already_exists"
	ErrNotPYMWYMIUser    ErrorCode = "not_pymwymi_user"
	ErrChallengeNotFound ErrorCode = "challenge_not_found"
	ErrBadInput          ErrorCode = "bad_input"
	ErrVotingFinished    ErrorCode = "voting_finished"
	ErrNotParticipant    ErrorCode = "not_participant"
	ErrInternal          ErrorCode = "internal_error"
)

// Error represents an application-specific error. Application errors can be
// unwrapped by the caller to extract out the code & message.

type Error struct {
	// Machine-readable error code.
	Code ErrorCode

	// Human-readable error message.
	Message string
}

// Error implements the error interface. Not used by the application otherwise.
func (e *Error) Error() string {
	return fmt.Sprintf("pymwymi error: code=%s message=%s", e.Code, e.Message)
}

// ErrorCode unwraps an application error and returns its code.
// Non-application errors always return EINTERNAL.
func GetErrorCode(err error) ErrorCode {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) && e != nil {
		return e.Code
	}
	return ErrInternal
}

// ErrorMessage unwraps an application error and returns its message.
// Non-application errors always return "Internal error".
func ErrorMessage(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}
	return err.Error()
}

// Errorf is a helper function to return an Error with a given code and formatted message.
func Errorf(code ErrorCode, format string, args ...any) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
