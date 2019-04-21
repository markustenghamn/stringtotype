package stringtotype

import (
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	kind := reflect.Bool

	v := "true"
	compareTypes(v, kind, t)

	v = "false"
	compareTypes(v, kind, t)

	v = "0"
	compareTypes(v, kind, t)

	v = "1"
	compareTypes(v, kind, t)

	v = "1.0"
	compareInvalidTypes(v, kind, t)

	v = "a"
	compareInvalidTypes(v, kind, t)

	v = "first"
	compareInvalidTypes(v, kind, t)

	v = "2"
	compareInvalidTypes(v, kind, t)

	v = "not true"
	compareInvalidTypes(v, kind, t)
}

func TestInt(t *testing.T) {
	kind := reflect.Int

	v := "12"
	compareTypes(v, kind, t)

	v = "200"
	compareTypes(v, kind, t)

	v = "89832983283"
	compareTypes(v, kind, t)

	v = "-12"
	compareTypes(v, kind, t)

	v = "1.0"
	compareInvalidTypes(v, kind, t)

	v = "a"
	compareInvalidTypes(v, kind, t)

	v = "first"
	compareInvalidTypes(v, kind, t)

	v = "0.2122"
	compareInvalidTypes(v, kind, t)

	v = "not true"
	compareInvalidTypes(v, kind, t)
}

func TestFloat64(t *testing.T) {
	kind := reflect.Float64

	v := "0.21312"
	compareTypes(v, kind, t)

	v = "43434.00392"
	compareTypes(v, kind, t)

	v = "1.0"
	compareTypes(v, kind, t)

	v = "0.0"
	compareTypes(v, kind, t)

	v = "1.0x"
	compareInvalidTypes(v, kind, t)

	v = "a"
	compareInvalidTypes(v, kind, t)

	v = "first"
	compareInvalidTypes(v, kind, t)

	v = "2"
	compareInvalidTypes(v, kind, t)

	v = "not true"
	compareInvalidTypes(v, kind, t)
}

func compareTypes(v string, kind reflect.Kind, t *testing.T) {
	if x := GetType(v); x.Kind() != kind {
		t.Error("Expected ", kind.String(), ", got ", x.Kind())
	}
}

func compareInvalidTypes(v string, kind reflect.Kind, t *testing.T) {
	if x := GetType(v); x.Kind() == kind {
		t.Error("Expected to not get ", kind.String(), ", got ", x.Kind())
	}
}
