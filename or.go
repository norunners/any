// Package any is an empty interface replacement.
package any

// Or sets the Value to the default when not OK and returns the Value.
// Otherwise, the original Value is returned and the default is ignored.
// This is useful for providing a default before using the underlying value.
func (v Value) Or(i interface{}) Value {
	if !v.OK() {
		v.i = i
	}
	return v
}

// BoolOr returns the value as a bool type or the default when the Value is not of type bool
func (v Value) BoolOr(d bool) bool {
	b, ok := v.i.(bool)
	if !ok {
		return d
	}
	return b
}

// IntOr returns the value as a int type or the default when the Value is not of type int.
func (v Value) IntOr(d int) int {
	i, ok := v.i.(int)
	if !ok {
		return d
	}
	return i
}

// Int8Or returns the value as a int8 type or the default when the Value is not of type int8.
func (v Value) Int8Or(d int8) int8 {
	i, ok := v.i.(int8)
	if !ok {
		return d
	}
	return i
}

// Int16Or returns the value as a int16 type or the default when the Value is not of type int16.
func (v Value) Int16Or(d int16) int16 {
	i, ok := v.i.(int16)
	if !ok {
		return d
	}
	return i
}

// Int32Or returns the value as a int32 type or the default when the Value is not of type int32.
func (v Value) Int32Or(d int32) int32 {
	i, ok := v.i.(int32)
	if !ok {
		return d
	}
	return i
}

// Int64Or returns the value as a int64 type or the default when the Value is not of type int64.
func (v Value) Int64Or(d int64) int64 {
	i, ok := v.i.(int64)
	if !ok {
		return d
	}
	return i
}

// UintOr returns the value as a uint type or the default when the Value is not of type uint.
func (v Value) UintOr(d uint) uint {
	i, ok := v.i.(uint)
	if !ok {
		return d
	}
	return i
}

// Uint8Or returns the value as a uint8 type or the default when the Value is not of type uint8.
func (v Value) Uint8Or(d uint8) uint8 {
	i, ok := v.i.(uint8)
	if !ok {
		return d
	}
	return i
}

// Uint16Or returns the value as a uint16 type or the default when the Value is not of type uint16.
func (v Value) Uint16Or(d uint16) uint16 {
	i, ok := v.i.(uint16)
	if !ok {
		return d
	}
	return i
}

// Uint32Or returns the value as a uint32 type or the default when the Value is not of type uint32.
func (v Value) Uint32Or(d uint32) uint32 {
	i, ok := v.i.(uint32)
	if !ok {
		return d
	}
	return i
}

// Uint64Or returns the value as a uint64 type or the default when the Value is not of type uint64.
func (v Value) Uint64Or(d uint64) uint64 {
	i, ok := v.i.(uint64)
	if !ok {
		return d
	}
	return i
}

// UintptrOr returns the value as a uintptr type or the default when the Value is not of type uintptr.
func (v Value) UintptrOr(d uintptr) uintptr {
	i, ok := v.i.(uintptr)
	if !ok {
		return d
	}
	return i
}

// Float32Or returns the value as a float32 type or the default when the Value is not of type float32.
func (v Value) Float32Or(d float32) float32 {
	f, ok := v.i.(float32)
	if !ok {
		return d
	}
	return f
}

// Float64Or returns the value as a float64 type or the default when the Value is not of type float64.
func (v Value) Float64Or(d float64) float64 {
	f, ok := v.i.(float64)
	if !ok {
		return d
	}
	return f
}

// Complex64Or returns the value as a complex64 type or the default when the Value is not of type complex64.
func (v Value) Complex64Or(d complex64) complex64 {
	c, ok := v.i.(complex64)
	if !ok {
		return d
	}
	return c
}

// Complex128Or returns the value as a complex128 type or the default when the Value is not of type complex128.
func (v Value) Complex128Or(d complex128) complex128 {
	c, ok := v.i.(complex128)
	if !ok {
		return d
	}
	return c
}

// ByteOr returns the value as a byte type or the default when the Value is not of type byte.
func (v Value) ByteOr(d byte) byte {
	b, ok := v.i.(byte)
	if !ok {
		return d
	}
	return b
}

// BytesOr returns the value as a []byte type or the default when the Value is not of type []byte.
func (v Value) BytesOr(d []byte) []byte {
	b, ok := v.i.([]byte)
	if !ok {
		return d
	}
	return b
}

// RuneOr returns the value as a rune type or the default when the Value is not of type rune.
func (v Value) RuneOr(d rune) rune {
	r, ok := v.i.(rune)
	if !ok {
		return d
	}
	return r
}

// StringOr returns the value as a string type or the default when the Value is not of type string.
func (v Value) StringOr(d string) string {
	s, ok := v.i.(string)
	if !ok {
		return d
	}
	return s
}

// InterfaceOr provides the underlying value as an empty interface or the default when not OK.
func (v Value) InterfaceOr(d interface{}) interface{} {
	if !v.OK() {
		return d
	}
	return v.i
}
