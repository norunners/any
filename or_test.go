package any_test

import (
	"reflect"
	"testing"

	"github.com/norunners/any"
)

func TestValue_Or(t *testing.T) {
	tests := []struct {
		name     string
		value    any.Value
		def      string
		expected string
	}{
		{
			name: "zero string from zero to default",
		},
		{
			name:     "string from zero to default",
			def:      "world",
			expected: "world",
		},
		{
			name:  "zero string from value",
			value: any.ValueOf(""),
			def:   "world",
		},
		{
			name:     "string from value",
			value:    any.ValueOf("hello"),
			def:      "world",
			expected: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.value.Or(test.def).String()
			if actual != test.expected {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_TypeOr(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "BoolOr",
			expected: true,
		},
		{
			method:   "IntOr",
			expected: 42,
		},
		{
			method:   "Int8Or",
			expected: int8(42),
		},
		{
			method:   "Int16Or",
			expected: int16(42),
		},
		{
			method:   "Int32Or",
			expected: int32(42),
		},
		{
			method:   "Int64Or",
			expected: int64(42),
		},
		{
			method:   "UintOr",
			expected: uint(42),
		},
		{
			method:   "Uint8Or",
			expected: uint8(42),
		},
		{
			method:   "Uint16Or",
			expected: uint16(42),
		},
		{
			method:   "Uint32Or",
			expected: uint32(42),
		},
		{
			method:   "Uint64Or",
			expected: uint64(42),
		},
		{
			method:   "UintptrOr",
			expected: uintptr(42),
		},
		{
			method:   "Float32Or",
			expected: float32(42),
		},
		{
			method:   "Float64Or",
			expected: 42.0,
		},
		{
			method:   "Complex64Or",
			expected: complex64(42),
		},
		{
			method:   "Complex128Or",
			expected: 42i,
		},
		{
			method:   "ByteOr",
			expected: byte(42),
		},
		{
			method:   "BytesOr",
			expected: []byte("hello"),
		},
		{
			method:   "RuneOr",
			expected: 'h',
		},
		{
			method:   "StringOr",
			expected: "hello",
		},
		{
			method:   "InterfaceOr",
			expected: pairFull,
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			value := any.ValueOf(test.expected)
			rv := reflect.ValueOf(value)
			rm := rv.MethodByName(test.method)
			d := reflect.ValueOf(test.expected)
			in := []reflect.Value{reflect.New(d.Type()).Elem()}
			actual := rm.Call(in)[0].Interface()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_TypeOr_Zero(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "BoolOr",
			expected: true,
		},
		{
			method:   "IntOr",
			expected: 42,
		},
		{
			method:   "Int8Or",
			expected: int8(42),
		},
		{
			method:   "Int16Or",
			expected: int16(42),
		},
		{
			method:   "Int32Or",
			expected: int32(42),
		},
		{
			method:   "Int64Or",
			expected: int64(42),
		},
		{
			method:   "UintOr",
			expected: uint(42),
		},
		{
			method:   "Uint8Or",
			expected: uint8(42),
		},
		{
			method:   "Uint16Or",
			expected: uint16(42),
		},
		{
			method:   "Uint32Or",
			expected: uint32(42),
		},
		{
			method:   "Uint64Or",
			expected: uint64(42),
		},
		{
			method:   "UintptrOr",
			expected: uintptr(42),
		},
		{
			method:   "Float32Or",
			expected: float32(42),
		},
		{
			method:   "Float64Or",
			expected: 42.0,
		},
		{
			method:   "Complex64Or",
			expected: complex64(42),
		},
		{
			method:   "Complex128Or",
			expected: 42i,
		},
		{
			method:   "ByteOr",
			expected: byte(42),
		},
		{
			method:   "BytesOr",
			expected: []byte("hello"),
		},
		{
			method:   "RuneOr",
			expected: 'h',
		},
		{
			method:   "StringOr",
			expected: "hello",
		},
		{
			method:   "InterfaceOr",
			expected: pairFull,
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			rv := reflect.ValueOf(any.Value{})
			rm := rv.MethodByName(test.method)
			d := reflect.ValueOf(test.expected)
			in := []reflect.Value{d}
			actual := rm.Call(in)[0].Interface()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}
