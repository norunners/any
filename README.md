# any

[![Go Reference](https://pkg.go.dev/badge/github.com/norunners/any.svg)](https://pkg.go.dev/github.com/norunners/any)

Package `any` is an empty interface replacement.

# Install
```
go get github.com/norunners/any
```

# Examples

### ValueOf
`ValueOf` creates a `Value` from any value.
```go
v := any.ValueOf("hello")
fmt.Println(v.String())
// "hello"
```

### OK
`OK` determines whether `Value` is not the zero `Value`.
```go
if v := any.ValueOf("hello"); v.OK() {
	fmt.Println(v.String())
	// "hello"
}
```

### Default
`Default` creates a new `Value` from the default when not `OK`.
```go
var v any.Value
v = v.Default("world")
fmt.Println(v.String())
// "world"
```

### Set
`Set` assigns `Value` to a non-nil pointer.
```go
v := any.ValueOf(pair{q: "meaning", a: 42})
p := pair{}
if err := v.Set(&p); err != nil {
	// Handle error.
}
fmt.Println(p)
// {"meaning", 42}
```

### Interface
`Interface` provides the underlying value as an empty interface.
```go
v := any.ValueOf(pair{q: "meaning", a: 42})
if p, ok := v.Interface().(pair); ok {
	fmt.Println(p)
	// {"meaning", 42}
}
```

### Map
`Map` represents a map of any values.
```go
m := any.Map{"meaning": any.ValueOf(42)}
v := m["meaning"]
fmt.Println(v.Int())
// 42
```

### MapOf
`MapOf` makes a `Map` from a map of any values.
```go
m := any.MapOf(map[string]interface{}{"meaning": 42})
v := m["meaning"]
fmt.Println(v.Int())
// 42
```

### Put
`Put` puts the key value into the returned `Map`.
```go
var m any.Map
m = m.Put("meaning", 42)
v := m["meaning"]
fmt.Println(v.Int())
```

# License
* [MIT License](LICENSE)
