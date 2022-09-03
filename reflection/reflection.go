package main

import (
	"fmt"
	"reflect"
)

func walk(x any, fn func(string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
}

func getValue(x any) reflect.Value {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return v
}

type T struct {
	A string
	B int
	C string
}

func f(s string) {
	fmt.Println(s)
}

func main() {
	t := T{
		A: "aaaaa",
		B: 11,
		C: "ccccc",
	}
	walk(t, f)
	walk(&t, f)

}
