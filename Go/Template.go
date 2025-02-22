package main

import (
	"fmt"
	"reflect"
)

func templateMethod() string {
	return "Template"
}

func main() {
	// Test cases go here
	fmt.Println(reflect.DeepEqual(templateMethod(), "Template"))
}
