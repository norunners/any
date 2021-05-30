package any

// Map represents a map of any values.
type Map map[string]Value

// MapOf makes a Map from a map of any values.
// This is useful for handling empty interface maps.
// The provided Map is always safe to use and never nil.
// The MapOf nil is equivalent to an empty Map, e.g. any.Map{}.
func MapOf(m map[string]interface{}) Map {
	vm := make(Map, len(m))
	for k, v := range m {
		vm[k] = ValueOf(v)
	}
	return vm
}

// Put puts the key value into the returned Map.
// This is useful for handling empty interface values.
// A new Map is returned in the case of using a nil Map.
func (m Map) Put(k string, v interface{}) Map {
	if m == nil {
		m = make(Map)
	}
	m[k] = ValueOf(v)
	return m
}
