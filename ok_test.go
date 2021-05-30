package any_test

import (
	"reflect"
	"testing"

	"github.com/norunners/any"
)

func TestValue_OK(t *testing.T) {
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
			actual := test.value.OK()
			if actual != test.expected {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_TypeOK(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "BoolOK",
			expected: true,
		},
		{
			method:   "IntOK",
			expected: 42,
		},
		{
			method:   "Int8OK",
			expected: int8(42),
		},
		{
			method:   "Int16OK",
			expected: int16(42),
		},
		{
			method:   "Int32OK",
			expected: int32(42),
		},
		{
			method:   "Int64OK",
			expected: int64(42),
		},
		{
			method:   "UintOK",
			expected: uint(42),
		},
		{
			method:   "Uint8OK",
			expected: uint8(42),
		},
		{
			method:   "Uint16OK",
			expected: uint16(42),
		},
		{
			method:   "Uint32OK",
			expected: uint32(42),
		},
		{
			method:   "Uint64OK",
			expected: uint64(42),
		},
		{
			method:   "UintptrOK",
			expected: uintptr(42),
		},
		{
			method:   "Float32OK",
			expected: float32(42),
		},
		{
			method:   "Float64OK",
			expected: 42.0,
		},
		{
			method:   "Complex64OK",
			expected: complex64(42),
		},
		{
			method:   "Complex128OK",
			expected: 42i,
		},
		{
			method:   "ByteOK",
			expected: byte(42),
		},
		{
			method:   "BytesOK",
			expected: []byte("hello"),
		},
		{
			method:   "RuneOK",
			expected: 'h',
		},
		{
			method:   "StringOK",
			expected: "hello",
		},
		{
			method:   "InterfaceOK",
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

func TestValue_TypeOK_Zero(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "BoolOK",
			expected: false,
		},
		{
			method:   "IntOK",
			expected: 0,
		},
		{
			method:   "Int8OK",
			expected: int8(0),
		},
		{
			method:   "Int16OK",
			expected: int16(0),
		},
		{
			method:   "Int32OK",
			expected: int32(0),
		},
		{
			method:   "Int64OK",
			expected: int64(0),
		},
		{
			method:   "UintOK",
			expected: uint(0),
		},
		{
			method:   "Uint8OK",
			expected: uint8(0),
		},
		{
			method:   "Uint16OK",
			expected: uint16(0),
		},
		{
			method:   "Uint32OK",
			expected: uint32(0),
		},
		{
			method:   "Uint64OK",
			expected: uint64(0),
		},
		{
			method:   "UintptrOK",
			expected: uintptr(0),
		},
		{
			method:   "Float32OK",
			expected: float32(0),
		},
		{
			method:   "Float64OK",
			expected: 0.0,
		},
		{
			method:   "Complex64OK",
			expected: complex64(0),
		},
		{
			method:   "Complex128OK",
			expected: 0i,
		},
		{
			method:   "ByteOK",
			expected: byte(0),
		},
		{
			method:   "BytesOK",
			expected: []byte(nil),
		},
		{
			method:   "RuneOK",
			expected: '\x00',
		},
		{
			method:   "StringOK",
			expected: "",
		},
		{
			method: "InterfaceOK",
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
