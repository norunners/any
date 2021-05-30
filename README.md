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

### Ok
`Ok` determines whether `Value` is not the zero `Value`.
```go
v := any.ValueOf("hello")
fmt.Println(v.Ok())
// true
```

### [Type]Ok
[Type]Ok returns the value as a [Type] type and a bool whether the Value is of type [Type].
```go
v := any.ValueOf("hello")
s, ok := v.StringOk()
fmt.Println(s, ok)
// "hello" true
```

### Or
`Or` sets the Value to the default when not `Ok` and returns the `Value`.
```go
var v any.Value
v = v.Or("world")
fmt.Println(v.String())
// "world"
```
### [Type]Or
[Type]Or returns the value as a [Type] type or the default when the Value is not of type [Type].
```go
var v any.Value
s := v.StringOr("world")
fmt.Println(s)
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
// 42
```

# License
* [MIT License](LICENSE)
