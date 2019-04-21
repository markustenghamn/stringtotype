package stringtotype

import (
	"reflect"
	"strconv"
)

// GetType takes a string and tries to determine what type the string is
func GetType(s string) reflect.Type {
	//TODO should Array and Map be supported?

	// Bool takes true, false, 1 or 0
	b, err := strconv.ParseBool(s)
	if err == nil {
		return reflect.TypeOf(b)
	}

	// Int should convert whole numbers such as 1
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return reflect.TypeOf(int(i))
	}

	// Float should convert any string with a decimal 1.0
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return reflect.TypeOf(f)
	}

	// String is a fallback if we can not convert to anything else
	return reflect.TypeOf(s)
}
