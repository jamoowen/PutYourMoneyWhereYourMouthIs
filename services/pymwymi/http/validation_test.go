package http_test

import (
	"testing"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/http"
)

func TestNewStringValidator(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fieldName  string
		value      string
		validators []func(string) error
		wantError  bool
	}{
		{
			name:       "transaction hash is valid",
			value:      "0xbac951afb9b2e756ab5aa8dc945c13661f6337146e5305a3cd0062f521de9239",
			fieldName:  "transactionHash",
			validators: []func(string) error{http.CheckMaxChars(66), http.CheckMinChars(66)},
			wantError:  false,
		},
		{
			name:       "too short",
			value:      "0x123",
			fieldName:  "transactionHash",
			validators: []func(string) error{http.CheckMaxChars(66), http.CheckMinChars(66)},
			wantError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := http.NewStringValidator(tt.fieldName, tt.value, tt.validators...).Validate()
			if got == nil && tt.wantError {
				t.Errorf("NewStringValidator() = %v, want %v", got, tt.wantError)
			} else if got != nil && !tt.wantError {
				t.Errorf("NewStringValidator() = %v, want %v", got, tt.wantError)
			}
		})
	}
}
