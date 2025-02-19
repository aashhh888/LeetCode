package main

import (
	"fmt"
	"reflect"
	"shared"
)

func addTwoNumbers(l1 *shared.ListNode, l2 *shared.ListNode, carry int) *shared.ListNode {
	l1Val := 0
	var l1Next *shared.ListNode
	l2Val := 0
	var l2Next *shared.ListNode

	if l1 != nil {
		l1Val = l1.Val
		l1Next = l1.Next
	}

	if l2 != nil {
		l2Val = l2.Val
		l2Next = l2.Next
	}

	c, r := divmod(l1Val+l2Val+carry, 10)

	if l1 != nil || l2 != nil {
		return &shared.ListNode{Val: r, Next: addTwoNumbers(l1Next, l2Next, c)}
	} else {
		if r == 0 {
			return nil
		} else {
			return &shared.ListNode{Val: r, Next: nil}
		}
	}
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func main() {
	fmt.Println(reflect.DeepEqual(addTwoNumbers(shared.BuildListNode([]int{2, 4, 3}), shared.BuildListNode([]int{5, 6, 4}), 0),
		shared.BuildListNode([]int{7, 0, 8})))
	fmt.Println(reflect.DeepEqual(addTwoNumbers(shared.BuildListNode([]int{0}), shared.BuildListNode([]int{0}), 0),
		shared.BuildListNode([]int{0})))
	fmt.Println(reflect.DeepEqual(addTwoNumbers(shared.BuildListNode([]int{9, 9, 9, 9, 9, 9, 9}), shared.BuildListNode([]int{9, 9, 9, 9}), 0),
		shared.BuildListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})))
}
