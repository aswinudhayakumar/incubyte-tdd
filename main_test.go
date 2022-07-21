package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsLeapYear(t *testing.T) {
	const neagativeYearMessgae = "Year cannot be nagative"
	tests := []struct {
		description          string
		year                 int
		expected             bool
		expectedError        error
		expectedErrorMessage string
	}{
		{
			description:   "Divisible by 400 and returns true",
			year:          2000,
			expected:      true,
			expectedError: nil,
		},
		{
			description:   "Divisible by 100 but not by 400 and returns false",
			year:          1700,
			expected:      false,
			expectedError: nil,
		},
		{
			description:   "Divisible by 4 but not by 100 and returns true",
			year:          2008,
			expected:      true,
			expectedError: nil,
		},
		{
			description:   "Not divisible by 4 and returns false",
			year:          2017,
			expected:      false,
			expectedError: nil,
		},
		{
			description:   "Divisible by 4000 and returns false",
			year:          4000,
			expected:      false,
			expectedError: nil,
		},
		{
			description:          "Negative year and returns false",
			year:                 -2000,
			expected:             false,
			expectedError:        NegativeYearError,
			expectedErrorMessage: neagativeYearMessgae,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual, err := IsLeapYear(test.year)
			assert.Equal(t, test.expectedError, err)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}
