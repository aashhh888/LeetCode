package main

import (
	"fmt"
	"reflect"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	loc := 0
	var res []string
	zig := true

	for _, c := range s {
		if len(res) > loc {
			res[loc] = res[loc] + string(c)
		} else {
			res = append(res, string(c))
		}

		if zig {
			loc = loc + 1
		} else {
			loc = loc - 1
		}

		if loc == 0 {
			zig = true
		} else if loc == (numRows - 1) {
			zig = false
		}
	}

	return strings.Join(res, "")
}

func main() {
	fmt.Println(reflect.DeepEqual(convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR"))
	fmt.Println(reflect.DeepEqual(convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI"))
	fmt.Println(reflect.DeepEqual(convert("A", 1), "A"))
	fmt.Println(reflect.DeepEqual(convert("AB", 1), "AB"))
}
