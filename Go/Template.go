package main

import (
	"fmt"
	"reflect"
)

func templateMethod() string {
	return "Template"
}

func main() {
	fmt.Println(reflect.DeepEqual(templateMethod(), "Template"))
}
