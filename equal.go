package any

import "reflect"

// Equaler provides a way to let an underlying value define equality for Value.
type Equaler interface {
	// Equal determines if a Value is equal to another Value.
	Equal(Value) bool
}

// Equal satisfies Equaler for Value.
// Equality will be determined by the underlying values when Equaler is satisfied by either value.
// Otherwise, equality is determined by reflect.DeepEqual on the underlying values.
func (v Value) Equal(other Value) bool {
	if eq, ok := v.i.(Equaler); ok {
		return eq.Equal(other)
	}
	if eq, ok := other.i.(Equaler); ok {
		return eq.Equal(v)
	}
	return reflect.DeepEqual(v.i, other.i)
}

var _ Equaler = (*Value)(nil)

// Equal determines if a Map is equal to another Map.
func (m Map) Equal(other Map) bool {
	if len(m) != len(other) {
		return false
	}
	for k, v := range m {
		if !v.Equal(other[k]) {
			return false
		}
	}
	return true
}
