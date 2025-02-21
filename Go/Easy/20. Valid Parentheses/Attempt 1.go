package main

import (
	"fmt"
	"reflect"
)

func isValid(s string) bool {
	parantheses := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	if len(s)%2 == 1 {
		return false
	} else {
		_, exists := parantheses[string(s[0])]
		if !exists {
			return false
		}
	}

	res := []string{}

	for _, c := range s {
		value, exists := parantheses[string(c)]
		if exists {
			res = append(res, value)
		} else if len(res) != 0 && res[len(res)-1] == string(c) {
			res = res[:len(res)-1]
		} else {
			res = append(res, "?")
		}
	}
	return len(res) == 0
}

func main() {
	fmt.Println(reflect.DeepEqual(isValid("()"), true))
	fmt.Println(reflect.DeepEqual(isValid("()[]{}"), true))
	fmt.Println(reflect.DeepEqual(isValid("(]"), false))
	fmt.Println(reflect.DeepEqual(isValid("([])"), true))
	fmt.Println(reflect.DeepEqual(isValid("]"), false))
	fmt.Println(reflect.DeepEqual(isValid("){"), false))
	fmt.Println(reflect.DeepEqual(isValid("([}}])"), false))
	fmt.Println(reflect.DeepEqual(isValid("(){}}{"), false))
}
