package any_test

import (
	"reflect"
	"testing"

	"github.com/norunners/any"
)

func TestValue_Ok(t *testing.T) {
	tests := []struct {
		name     string
		value    any.Value
		expected bool
	}{
		{
			name: "false from zero",
		},
		{
			name:     "true from zero int",
			value:    any.ValueOf(0),
			expected: true,
		},
		{
			name:     "true from int",
			value:    any.ValueOf(42),
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.value.Ok()
			if actual != test.expected {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_TypeOk(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "BoolOk",
			expected: true,
		},
		{
			method:   "IntOk",
			expected: 42,
		},
		{
			method:   "Int8Ok",
			expected: int8(42),
		},
		{
			method:   "Int16Ok",
			expected: int16(42),
		},
		{
			method:   "Int32Ok",
			expected: int32(42),
		},
		{
			method:   "Int64Ok",
			expected: int64(42),
		},
		{
			method:   "UintOk",
			expected: uint(42),
		},
		{
			method:   "Uint8Ok",
			expected: uint8(42),
		},
		{
			method:   "Uint16Ok",
			expected: uint16(42),
		},
		{
			method:   "Uint32Ok",
			expected: uint32(42),
		},
		{
			method:   "Uint64Ok",
			expected: uint64(42),
		},
		{
			method:   "UintptrOk",
			expected: uintptr(42),
		},
		{
			method:   "Float32Ok",
			expected: float32(42),
		},
		{
			method:   "Float64Ok",
			expected: 42.0,
		},
		{
			method:   "Complex64Ok",
			expected: complex64(42),
		},
		{
			method:   "Complex128Ok",
			expected: 42i,
		},
		{
			method:   "ByteOk",
			expected: byte(42),
		},
		{
			method:   "BytesOk",
			expected: []byte("hello"),
		},
		{
			method:   "RuneOk",
			expected: 'h',
		},
		{
			method:   "StringOk",
			expected: "hello",
		},
		{
			method:   "InterfaceOk",
			expected: pairFull,
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			value := any.ValueOf(test.expected)
			rv := reflect.ValueOf(value)
			rm := rv.MethodByName(test.method)
			out := rm.Call(nil)
			if ok := out[1].Interface().(bool); !ok {
				t.Fatalf("unexpected ok bool: %v", ok)
			}
			actual := out[0].Interface()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_TypeOk_Zero(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "BoolOk",
			expected: false,
		},
		{
			method:   "IntOk",
			expected: 0,
		},
		{
			method:   "Int8Ok",
			expected: int8(0),
		},
		{
			method:   "Int16Ok",
			expected: int16(0),
		},
		{
			method:   "Int32Ok",
			expected: int32(0),
		},
		{
			method:   "Int64Ok",
			expected: int64(0),
		},
		{
			method:   "UintOk",
			expected: uint(0),
		},
		{
			method:   "Uint8Ok",
			expected: uint8(0),
		},
		{
			method:   "Uint16Ok",
			expected: uint16(0),
		},
		{
			method:   "Uint32Ok",
			expected: uint32(0),
		},
		{
			method:   "Uint64Ok",
			expected: uint64(0),
		},
		{
			method:   "UintptrOk",
			expected: uintptr(0),
		},
		{
			method:   "Float32Ok",
			expected: float32(0),
		},
		{
			method:   "Float64Ok",
			expected: 0.0,
		},
		{
			method:   "Complex64Ok",
			expected: complex64(0),
		},
		{
			method:   "Complex128Ok",
			expected: 0i,
		},
		{
			method:   "ByteOk",
			expected: byte(0),
		},
		{
			method:   "BytesOk",
			expected: []byte(nil),
		},
		{
			method:   "RuneOk",
			expected: '\x00',
		},
		{
			method:   "StringOk",
			expected: "",
		},
		{
			method: "InterfaceOk",
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			rv := reflect.ValueOf(any.Value{})
			rm := rv.MethodByName(test.method)
			out := rm.Call(nil)
			if ok := out[1].Interface().(bool); ok {
				t.Fatalf("unexpected ok bool: %v", ok)
			}
			actual := out[0].Interface()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}
