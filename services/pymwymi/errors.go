package pymwymi

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code=%s message=%s", e.Code, e.Message)
}
