package any_test

import (
	"reflect"
	"testing"

	"github.com/norunners/any"
)

func TestMap_Empty(t *testing.T) {
	empty := any.Map{}
	if m := any.MapOf(nil); !reflect.DeepEqual(m, empty) {
		t.Errorf("expected equal empty Map: %+v with: %+v", m, empty)
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		expected map[string]interface{}
	}{
		{
			name:     "empty",
			expected: map[string]interface{}{},
		},
		{
			name: "single",
			expected: map[string]interface{}{
				"meaning": 42,
			},
		},
		{
			name: "double",
			expected: map[string]interface{}{
				"key":     true,
				"meaning": 42,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := any.MapOf(test.expected)
			actual := make(map[string]interface{}, len(m))
			for k, v := range m {
				actual[k] = v.Interface()
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %+v but found: %+v", test.expected, actual)
			}
		})
	}
}

func TestMap_Put(t *testing.T) {
	tests := []struct {
		name string
		m    any.Map
	}{
		{
			name: "nil",
		},
		{
			name: "empty",
			m:    any.Map{},
		},
		{
			name: "make",
			m:    make(any.Map),
		},
	}

	k := "meaning"
	v := 42
	expected := any.Map{k: any.ValueOf(v)}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.m.Put(k, v)
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("expected: %+v but found: %+v", expected, actual)
			}
		})
	}
}
