package any

// OK determines whether Value is not the zero Value.
// This is useful to check before using the underlying value.
// Consider using Default over OK.
func (v Value) OK() bool {
	return v.i != nil
}

// BoolOK returns the value as a bool type and a bool whether the Value is of type bool.
func (v Value) BoolOK() (bool, bool) {
	b, ok := v.i.(bool)
	return b, ok
}

// IntOK returns the value as a int type and a bool whether the Value is of type int.
func (v Value) IntOK() (int, bool) {
	i, ok := v.i.(int)
	return i, ok
}

// Int8OK returns the value as a int8 type and a bool whether the Value is of type int8.
func (v Value) Int8OK() (int8, bool) {
	i, ok := v.i.(int8)
	return i, ok
}

// Int16OK returns the value as a int16 type and a bool whether the Value is of type int16.
func (v Value) Int16OK() (int16, bool) {
	i, ok := v.i.(int16)
	return i, ok
}

// Int32OK returns the value as a int32 type and a bool whether the Value is of type int32.
func (v Value) Int32OK() (int32, bool) {
	i, ok := v.i.(int32)
	return i, ok
}

// Int64OK returns the value as a int64 type and a bool whether the Value is of type int64.
func (v Value) Int64OK() (int64, bool) {
	i, ok := v.i.(int64)
	return i, ok
}

// UintOK returns the value as a uint type and a bool whether the Value is of type uint.
func (v Value) UintOK() (uint, bool) {
	i, ok := v.i.(uint)
	return i, ok
}

// Uint8OK returns the value as a uint8 type and a bool whether the Value is of type uint8.
func (v Value) Uint8OK() (uint8, bool) {
	i, ok := v.i.(uint8)
	return i, ok
}

// Uint16OK returns the value as a uint16 type and a bool whether the Value is of type uint16.
func (v Value) Uint16OK() (uint16, bool) {
	i, ok := v.i.(uint16)
	return i, ok
}

// Uint32OK returns the value as a uint32 type and a bool whether the Value is of type uint32.
func (v Value) Uint32OK() (uint32, bool) {
	i, ok := v.i.(uint32)
	return i, ok
}

// Uint64OK returns the value as a uint64 type and a bool whether the Value is of type uint64.
func (v Value) Uint64OK() (uint64, bool) {
	i, ok := v.i.(uint64)
	return i, ok
}

// UintptrOK returns the value as a uintptr type and a bool whether the Value is of type uintptr.
func (v Value) UintptrOK() (uintptr, bool) {
	i, ok := v.i.(uintptr)
	return i, ok
}

// Float32OK returns the value as a float32 type and a bool whether the Value is of type float32.
func (v Value) Float32OK() (float32, bool) {
	f, ok := v.i.(float32)
	return f, ok
}

// Float64OK returns the value as a float64 type and a bool whether the Value is of type float64.
func (v Value) Float64OK() (float64, bool) {
	f, ok := v.i.(float64)
	return f, ok
}

// Complex64OK returns the value as a complex64 type and a bool whether the Value is of type complex64.
func (v Value) Complex64OK() (complex64, bool) {
	c, ok := v.i.(complex64)
	return c, ok
}

// Complex128OK returns the value as a complex128 type and a bool whether the Value is of type complex128.
func (v Value) Complex128OK() (complex128, bool) {
	c, ok := v.i.(complex128)
	return c, ok
}

// ByteOK returns the value as a byte type and a bool whether the Value is of type byte.
func (v Value) ByteOK() (byte, bool) {
	b, ok := v.i.(byte)
	return b, ok
}

// BytesOK returns the value as a []byte type and a bool whether the Value is of type []byte.
func (v Value) BytesOK() ([]byte, bool) {
	b, ok := v.i.([]byte)
	return b, ok
}

// RuneOK returns the value as a rune type and a bool whether the Value is of type rune.
func (v Value) RuneOK() (rune, bool) {
	r, ok := v.i.(rune)
	return r, ok
}

// StringOK returns the value as a string type and a bool whether the Value is of type string.
func (v Value) StringOK() (string, bool) {
	s, ok := v.i.(string)
	return s, ok
}

// InterfaceOK provides the underlying value as an empty interface and a bool of OK.
func (v Value) InterfaceOK() (interface{}, bool) {
	return v.i, v.OK()
}
