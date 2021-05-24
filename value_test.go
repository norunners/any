package any_test

import (
	"reflect"
	"testing"

	"github.com/norunners/any"
)

func TestValue_Zero(t *testing.T) {
	var zero any.Value
	if value := any.ValueOf(nil); value != zero {
		t.Errorf("expected equal zero Value: %v with: %v", zero, value)
	}
}

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

func TestValue_Default(t *testing.T) {
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
			actual := test.value.Default(test.def).String()
			if actual != test.expected {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_Type(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "Bool",
			expected: true,
		},
		{
			method:   "Int",
			expected: 42,
		},
		{
			method:   "Int8",
			expected: int8(42),
		},
		{
			method:   "Int16",
			expected: int16(42),
		},
		{
			method:   "Int32",
			expected: int32(42),
		},
		{
			method:   "Int64",
			expected: int64(42),
		},
		{
			method:   "Uint",
			expected: uint(42),
		},
		{
			method:   "Uint8",
			expected: uint8(42),
		},
		{
			method:   "Uint16",
			expected: uint16(42),
		},
		{
			method:   "Uint32",
			expected: uint32(42),
		},
		{
			method:   "Uint64",
			expected: uint64(42),
		},
		{
			method:   "Uintptr",
			expected: uintptr(42),
		},
		{
			method:   "Float32",
			expected: float32(42),
		},
		{
			method:   "Float64",
			expected: 42.0,
		},
		{
			method:   "Complex64",
			expected: complex64(42),
		},
		{
			method:   "Complex128",
			expected: 42i,
		},
		{
			method:   "Byte",
			expected: byte(42),
		},
		{
			method:   "Bytes",
			expected: []byte("hello"),
		},
		{
			method:   "Rune",
			expected: 'h',
		},
		{
			method:   "String",
			expected: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			value := any.ValueOf(test.expected)
			rv := reflect.ValueOf(value)
			rm := rv.MethodByName(test.method)
			actual := rm.Call(nil)[0].Interface()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_Type_Zero(t *testing.T) {
	tests := []struct {
		method   string
		expected interface{}
	}{
		{
			method:   "Bool",
			expected: false,
		},
		{
			method:   "Int",
			expected: 0,
		},
		{
			method:   "Int8",
			expected: int8(0),
		},
		{
			method:   "Int16",
			expected: int16(0),
		},
		{
			method:   "Int32",
			expected: int32(0),
		},
		{
			method:   "Int64",
			expected: int64(0),
		},
		{
			method:   "Uint",
			expected: uint(0),
		},
		{
			method:   "Uint8",
			expected: uint8(0),
		},
		{
			method:   "Uint16",
			expected: uint16(0),
		},
		{
			method:   "Uint32",
			expected: uint32(0),
		},
		{
			method:   "Uint64",
			expected: uint64(0),
		},
		{
			method:   "Uintptr",
			expected: uintptr(0),
		},
		{
			method:   "Float32",
			expected: float32(0),
		},
		{
			method:   "Float64",
			expected: 0.0,
		},
		{
			method:   "Complex64",
			expected: complex64(0),
		},
		{
			method:   "Complex128",
			expected: 0i,
		},
		{
			method:   "Byte",
			expected: byte(0),
		},
		{
			method:   "Bytes",
			expected: []byte(nil),
		},
		{
			method:   "Rune",
			expected: '\x00',
		},
		{
			method:   "String",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			rv := reflect.ValueOf(any.Value{})
			rm := rv.MethodByName(test.method)
			actual := rm.Call(nil)[0].Interface()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_Interface(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
	}{
		{
			name: "nil",
		},
		{
			name:     "struct",
			expected: pairFull,
		},
		{
			name:     "bool",
			expected: true,
		},
		{
			name:     "int",
			expected: 42,
		},
		{
			name:     "int8",
			expected: int8(42),
		},
		{
			name:     "int16",
			expected: int16(42),
		},
		{
			name:     "int32",
			expected: int32(42),
		},
		{
			name:     "int64",
			expected: int64(42),
		},
		{
			name:     "uint",
			expected: uint(42),
		},
		{
			name:     "uint8",
			expected: uint8(42),
		},
		{
			name:     "uint16",
			expected: uint16(42),
		},
		{
			name:     "uint32",
			expected: uint32(42),
		},
		{
			name:     "uint64",
			expected: uint64(42),
		},
		{
			name:     "uintptr",
			expected: uintptr(42),
		},
		{
			name:     "float32",
			expected: float32(42),
		},
		{
			name:     "float64",
			expected: 42.0,
		},
		{
			name:     "complex64",
			expected: complex64(42),
		},
		{
			name:     "complex128",
			expected: 42i,
		},
		{
			name:     "byte",
			expected: byte(42),
		},
		{
			name:     "[]byte",
			expected: []byte("hello"),
		},
		{
			name:     "rune",
			expected: 'h',
		},
		{
			name:     "string",
			expected: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := any.ValueOf(test.expected).Interface()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_Set(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
	}{
		{
			name:     "struct",
			expected: pairFull,
		},
		{
			name:     "pointer",
			expected: &pairFull,
		},
		{
			name:     "pointer nil",
			expected: (*pair)(nil),
		},
		{
			name:     "bool",
			expected: true,
		},
		{
			name:     "int",
			expected: 42,
		},
		{
			name:     "int8",
			expected: int8(42),
		},
		{
			name:     "int16",
			expected: int16(42),
		},
		{
			name:     "int32",
			expected: int32(42),
		},
		{
			name:     "int64",
			expected: int64(42),
		},
		{
			name:     "uint",
			expected: uint(42),
		},
		{
			name:     "uint8",
			expected: uint8(42),
		},
		{
			name:     "uint16",
			expected: uint16(42),
		},
		{
			name:     "uint32",
			expected: uint32(42),
		},
		{
			name:     "uint64",
			expected: uint64(42),
		},
		{
			name:     "uintptr",
			expected: uintptr(42),
		},
		{
			name:     "float32",
			expected: float32(42),
		},
		{
			name:     "float64",
			expected: 42.0,
		},
		{
			name:     "complex64",
			expected: complex64(42),
		},
		{
			name:     "complex128",
			expected: 42i,
		},
		{
			name:     "byte",
			expected: byte(42),
		},
		{
			name:     "[]byte",
			expected: []byte("hello"),
		},
		{
			name:     "rune",
			expected: 'h',
		},
		{
			name:     "string",
			expected: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := any.ValueOf(test.expected)
			var actual interface{}
			if err := value.Set(&actual); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v but found: %v", test.expected, actual)
			}
		})
	}
}

func TestValue_Set_FromZero(t *testing.T) {
	var zero any.Value
	if err := zero.Set(&pair{}); err == nil {
		t.Error("expected error")
	}
}

func TestValue_Set_ToNil(t *testing.T) {
	if err := any.ValueOf(pairFull).Set(nil); err == nil {
		t.Error("expected error")
	}
}

func TestValue_Set_ToNilPointer(t *testing.T) {
	if err := any.ValueOf(pairFull).Set((*pair)(nil)); err == nil {
		t.Error("expected error")
	}
}

func TestValue_Set_ToNonPointer(t *testing.T) {
	if err := any.ValueOf(pairFull).Set(pair{}); err == nil {
		t.Error("expected error")
	}
}

func TestValue_Set_ToNotAssignable(t *testing.T) {
	if err := any.ValueOf(&pairFull).Set(&pair{}); err == nil {
		t.Error("expected error")
	}
}

type pair struct {
	q string
	a int
}

var pairFull = pair{
	q: "meaning",
	a: 42,
}
