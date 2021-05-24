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
m := map[string]any.Value{"key": any.ValueOf("hello")}
if v := m["key"]; v.OK() {
	fmt.Println(v.String())
	// "hello"
}
```

### Default
`Default` creates a new `Value` from the default when not `OK`.
```go
m := map[string]any.Value{}
v := m["key"].Default("world")
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

# License
* [MIT License](LICENSE)
