package main

import (
	"fmt"
	"reflect"
	"shared"
)

func addTwoNumbers(l1 *shared.ListNode, l2 *shared.ListNode) *shared.ListNode {
	return &shared.ListNode{}
}

func main() {
	shared.BuildListNode([]int{1, 2, 4})
	fmt.Println(reflect.DeepEqual(addTwoNumbers(&shared.ListNode{}, &shared.ListNode{}), &shared.ListNode{}))
}
