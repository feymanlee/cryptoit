//
// Package cryptoit
// @Author: feymanlee@gmail.com
// @Description:
// @File:  utils_test.go
// @Date: 2023/4/23 13:00
//

package cryptoit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyToByte(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected []byte
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "string input",
			input:    "test_string",
			expected: []byte("test_string"),
		},
		{
			name:     "byte slice input",
			input:    []byte{0x41, 0x42, 0x43},
			expected: []byte{0x41, 0x42, 0x43},
		},
		{
			name:     "integer input",
			input:    123,
			expected: []byte("123"),
		},
		{
			name:     "float input",
			input:    1.23,
			expected: []byte("1.23"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := AnyToByte(tc.input)
			assert.Equal(t, tc.expected, output, "AnyToByte(%v) should return %v", tc.input, tc.expected)
		})
	}
}
