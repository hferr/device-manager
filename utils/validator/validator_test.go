package validator_test

import (
	"github/hferr/device-manager/utils/validator"
	"testing"
)

func TestErrResponse(t *testing.T) {
	var testCases = map[string]struct {
		name     string
		input    any
		expected string
	}{
		"required": {
			input: struct {
				Name string `json:"name" validate:"required"`
			}{},
			expected: "Name is required",
		},
		"oneof": {
			input: struct {
				Status string `json:"status" validate:"oneof=active inactive"`
			}{},
			expected: "Status must be one of: active inactive",
		},
	}

	v := validator.New()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := v.Struct(tc.input)
			if res := validator.ErrResponse(err); res == nil || len(res.Errors) != 1 {
				t.Fatalf("expected error, got nil")
			} else if res.Errors[0] != tc.expected {
				t.Fatalf("expected %s, got %s", tc.expected, res.Errors[0])
			}
		})
	}
}
