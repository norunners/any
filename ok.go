package any

// Ok determines whether Value is not the zero Value.
// This is useful to check before using the underlying value.
// Consider using Default over Ok.
func (v Value) Ok() bool {
	return v.i != nil
}

// BoolOk returns the value as a bool type and a bool whether the Value is of type bool.
func (v Value) BoolOk() (bool, bool) {
	b, ok := v.i.(bool)
	return b, ok
}

// IntOk returns the value as a int type and a bool whether the Value is of type int.
func (v Value) IntOk() (int, bool) {
	i, ok := v.i.(int)
	return i, ok
}

// Int8Ok returns the value as a int8 type and a bool whether the Value is of type int8.
func (v Value) Int8Ok() (int8, bool) {
	i, ok := v.i.(int8)
	return i, ok
}

// Int16Ok returns the value as a int16 type and a bool whether the Value is of type int16.
func (v Value) Int16Ok() (int16, bool) {
	i, ok := v.i.(int16)
	return i, ok
}

// Int32Ok returns the value as a int32 type and a bool whether the Value is of type int32.
func (v Value) Int32Ok() (int32, bool) {
	i, ok := v.i.(int32)
	return i, ok
}

// Int64Ok returns the value as a int64 type and a bool whether the Value is of type int64.
func (v Value) Int64Ok() (int64, bool) {
	i, ok := v.i.(int64)
	return i, ok
}

// UintOk returns the value as a uint type and a bool whether the Value is of type uint.
func (v Value) UintOk() (uint, bool) {
	i, ok := v.i.(uint)
	return i, ok
}

// Uint8Ok returns the value as a uint8 type and a bool whether the Value is of type uint8.
func (v Value) Uint8Ok() (uint8, bool) {
	i, ok := v.i.(uint8)
	return i, ok
}

// Uint16Ok returns the value as a uint16 type and a bool whether the Value is of type uint16.
func (v Value) Uint16Ok() (uint16, bool) {
	i, ok := v.i.(uint16)
	return i, ok
}

// Uint32Ok returns the value as a uint32 type and a bool whether the Value is of type uint32.
func (v Value) Uint32Ok() (uint32, bool) {
	i, ok := v.i.(uint32)
	return i, ok
}

// Uint64Ok returns the value as a uint64 type and a bool whether the Value is of type uint64.
func (v Value) Uint64Ok() (uint64, bool) {
	i, ok := v.i.(uint64)
	return i, ok
}

// UintptrOk returns the value as a uintptr type and a bool whether the Value is of type uintptr.
func (v Value) UintptrOk() (uintptr, bool) {
	i, ok := v.i.(uintptr)
	return i, ok
}

// Float32Ok returns the value as a float32 type and a bool whether the Value is of type float32.
func (v Value) Float32Ok() (float32, bool) {
	f, ok := v.i.(float32)
	return f, ok
}

// Float64Ok returns the value as a float64 type and a bool whether the Value is of type float64.
func (v Value) Float64Ok() (float64, bool) {
	f, ok := v.i.(float64)
	return f, ok
}

// Complex64Ok returns the value as a complex64 type and a bool whether the Value is of type complex64.
func (v Value) Complex64Ok() (complex64, bool) {
	c, ok := v.i.(complex64)
	return c, ok
}

// Complex128Ok returns the value as a complex128 type and a bool whether the Value is of type complex128.
func (v Value) Complex128Ok() (complex128, bool) {
	c, ok := v.i.(complex128)
	return c, ok
}

// ByteOk returns the value as a byte type and a bool whether the Value is of type byte.
func (v Value) ByteOk() (byte, bool) {
	b, ok := v.i.(byte)
	return b, ok
}

// BytesOk returns the value as a []byte type and a bool whether the Value is of type []byte.
func (v Value) BytesOk() ([]byte, bool) {
	b, ok := v.i.([]byte)
	return b, ok
}

// RuneOk returns the value as a rune type and a bool whether the Value is of type rune.
func (v Value) RuneOk() (rune, bool) {
	r, ok := v.i.(rune)
	return r, ok
}

// StringOk returns the value as a string type and a bool whether the Value is of type string.
func (v Value) StringOk() (string, bool) {
	s, ok := v.i.(string)
	return s, ok
}

// InterfaceOk provides the underlying value as an empty interface and a bool of Ok.
func (v Value) InterfaceOk() (interface{}, bool) {
	return v.i, v.Ok()
}
