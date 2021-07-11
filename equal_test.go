package any_test

import (
	"testing"

	"github.com/norunners/any"
)

func TestValue_Equal(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		other    interface{}
		expected bool
	}{
		{
			name:     "zero",
			expected: true,
		},
		{
			name:     "false both bool",
			value:    false,
			other:    false,
			expected: true,
		},
		{
			name:     "false value bool",
			value:    false,
			other:    true,
			expected: false,
		},
		{
			name:     "false other bool",
			value:    true,
			other:    false,
			expected: false,
		},
		{
			name:     "true both bool",
			value:    true,
			other:    true,
			expected: true,
		},
		{
			name:     "false value equaler",
			value:    equaler{eq: false},
			expected: false,
		},
		{
			name:     "false other equaler",
			other:    equaler{eq: false},
			expected: false,
		},
		{
			name:     "true value equaler",
			value:    equaler{eq: true},
			expected: true,
		},
		{
			name:     "true other equaler",
			other:    equaler{eq: true},
			expected: true,
		},
		{
			name:     "false both equaler",
			value:    equaler{eq: false},
			other:    equaler{eq: false},
			expected: false,
		},
		{
			name:     "true mix equaler",
			value:    equaler{eq: true},
			other:    equaler{eq: false},
			expected: true,
		},
		{
			name:     "false mix equaler",
			value:    equaler{eq: false},
			other:    equaler{eq: true},
			expected: false,
		},
		{
			name:     "true both equaler",
			value:    equaler{eq: true},
			other:    equaler{eq: true},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := any.ValueOf(test.value)
			other := any.ValueOf(test.other)

			if actual := value.Equal(other); actual != test.expected {
				t.Errorf("expected: %t but found: %t", test.expected, actual)
			}
		})
	}
}

func TestMap_Equal(t *testing.T) {
	tests := []struct {
		name     string
		m        any.Map
		other    any.Map
		expected bool
	}{
		{
			name:     "zero",
			expected: true,
		},
		{
			name:     "empty",
			m:        any.MapOf(nil),
			expected: true,
		},
		{
			name:     "unequal length",
			m:        any.Map{}.Put("hello", "world"),
			expected: false,
		},
		{
			name:     "unequal types",
			m:        any.Map{}.Put("hello", "world"),
			other:    any.Map{}.Put("hello", 42),
			expected: false,
		},
		{
			name:     "unequal",
			m:        any.Map{}.Put("hello", "world"),
			other:    any.Map{}.Put("hello", "planet"),
			expected: false,
		},
		{
			name:     "equal",
			m:        any.Map{}.Put("hello", "world"),
			other:    any.Map{}.Put("hello", "world"),
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if actual := test.m.Equal(test.other); actual != test.expected {
				t.Errorf("expected: %t but found: %t", test.expected, actual)
			}
			if actual := test.other.Equal(test.m); actual != test.expected {
				t.Errorf("expected: %t but found: %t", test.expected, actual)
			}
		})
	}
}

type equaler struct {
	eq bool
}

func (e equaler) Equal(any.Value) bool {
	return e.eq
}

var _ any.Equaler = (*equaler)(nil)
