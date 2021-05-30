// Package any is an empty interface replacement.
package any

import (
	"fmt"
	"reflect"
)

// Value represents any value.
type Value struct {
	i interface{}
}

// ValueOf creates a Value from any value.
// The ValueOf nil is equivalent to the zero Value.
func ValueOf(i interface{}) Value {
	return Value{i: i}
}

// Bool returns the value as a bool type.
// The zero value of bool is returned when the value can not be type-asserted to a bool.
func (v Value) Bool() bool {
	b, _ := v.i.(bool)
	return b
}

// Int returns the value as a int type.
// The zero value of int is returned when the value can not be type-asserted to a int.
func (v Value) Int() int {
	i, _ := v.i.(int)
	return i
}

// Int8 returns the value as a int8 type.
// The zero value of int8 is returned when the value can not be type-asserted to a int8.
func (v Value) Int8() int8 {
	i, _ := v.i.(int8)
	return i
}

// Int16 returns the value as a int16 type.
// The zero value of int16 is returned when the value can not be type-asserted to a int16.
func (v Value) Int16() int16 {
	i, _ := v.i.(int16)
	return i
}

// Int32 returns the value as a int32 type.
// The zero value of int32 is returned when the value can not be type-asserted to a int32.
func (v Value) Int32() int32 {
	i, _ := v.i.(int32)
	return i
}

// Int64 returns the value as a int64 type.
// The zero value of int64 is returned when the value can not be type-asserted to a int64.
func (v Value) Int64() int64 {
	i, _ := v.i.(int64)
	return i
}

// Uint returns the value as a uint type.
// The zero value of uint is returned when the value can not be type-asserted to a uint.
func (v Value) Uint() uint {
	i, _ := v.i.(uint)
	return i
}

// Uint8 returns the value as a uint8 type.
// The zero value of uint8 is returned when the value can not be type-asserted to a uint8.
func (v Value) Uint8() uint8 {
	i, _ := v.i.(uint8)
	return i
}

// Uint16 returns the value as a uint16 type.
// The zero value of uint16 is returned when the value can not be type-asserted to a uint16.
func (v Value) Uint16() uint16 {
	i, _ := v.i.(uint16)
	return i
}

// Uint32 returns the value as a uint32 type.
// The zero value of uint32 is returned when the value can not be type-asserted to a uint32.
func (v Value) Uint32() uint32 {
	i, _ := v.i.(uint32)
	return i
}

// Uint64 returns the value as a uint64 type.
// The zero value of uint64 is returned when the value can not be type-asserted to a uint64.
func (v Value) Uint64() uint64 {
	i, _ := v.i.(uint64)
	return i
}

// Uintptr returns the value as a uintptr type.
// The zero value of uintptr is returned when the value can not be type-asserted to a uintptr.
func (v Value) Uintptr() uintptr {
	i, _ := v.i.(uintptr)
	return i
}

// Float32 returns the value as a float32 type.
// The zero value of float32 is returned when the value can not be type-asserted to a float32.
func (v Value) Float32() float32 {
	f, _ := v.i.(float32)
	return f
}

// Float64 returns the value as a float64 type.
// The zero value of float64 is returned when the value can not be type-asserted to a float64.
func (v Value) Float64() float64 {
	f, _ := v.i.(float64)
	return f
}

// Complex64 returns the value as a complex64 type.
// The zero value of complex64 is returned when the value can not be type-asserted to a complex64.
func (v Value) Complex64() complex64 {
	c, _ := v.i.(complex64)
	return c
}

// Complex128 returns the value as a complex128 type.
// The zero value of complex128 is returned when the value can not be type-asserted to a complex128.
func (v Value) Complex128() complex128 {
	c, _ := v.i.(complex128)
	return c
}

// Byte returns the value as a byte type.
// The zero value of byte is returned when the value can not be type-asserted to a byte.
func (v Value) Byte() byte {
	b, _ := v.i.(byte)
	return b
}

// Bytes returns the value as a []byte type.
// A nil []byte is returned when the value can not be type-asserted to a []byte.
func (v Value) Bytes() []byte {
	b, _ := v.i.([]byte)
	return b
}

// Rune returns the value as a rune type.
// The zero value of rune is returned when the value can not be type-asserted to a rune.
func (v Value) Rune() rune {
	r, _ := v.i.(rune)
	return r
}

// String returns the value as a string type.
// The zero value of string is returned when the value can not be type-asserted to a string.
func (v Value) String() string {
	s, _ := v.i.(string)
	return s
}

// Interface provides the underlying value as an empty interface.
// This is useful for type-asserting to a concrete type.
// Consider using Set over Interface.
// A nil value is provided for the zero Value.
func (v Value) Interface() interface{} {
	return v.i
}

// Set assigns Value to a non-nil pointer.
// An error is returned when the assignment fails.
// A Value can not be Set when not OK.
// A Value can not be Set to a non-pointer type nor a nil pointer.
// A Value can not be Set to another type.
// All panics are unexpected.
func (v Value) Set(ptr interface{}) error {
	if !v.OK() {
		return fmt.Errorf("any: Value must be OK")
	}
	dst := reflect.ValueOf(ptr)
	if k := dst.Kind(); k != reflect.Ptr || dst.IsNil() {
		return fmt.Errorf("any: ptr must be a non-nil pointer: %v", k)
	}
	src := reflect.ValueOf(v.i)
	dst = dst.Elem()
	if st, dt := src.Type(), dst.Type(); !st.AssignableTo(dt) {
		return fmt.Errorf("any: unassignable from: %v to: %v", st, dt)
	}
	dst.Set(src)
	return nil
}
