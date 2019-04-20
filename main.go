package stringtotype

import (
	"reflect"
	"strconv"
)

func GetType(s string) reflect.Type {
	b, err := strconv.ParseBool(s)
	if err == nil {
		return reflect.TypeOf(b)
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return reflect.TypeOf(i)
	}

	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return reflect.TypeOf(f)
	}

	return reflect.TypeOf(s)
}
