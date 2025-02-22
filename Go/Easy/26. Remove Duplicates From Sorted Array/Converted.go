package main

import (
	"fmt"
	"reflect"
)

func removeDuplicates(nums []int) int {
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}

func main() {
	case1 := []int{1, 1, 2}
	fmt.Println(reflect.DeepEqual(removeDuplicates(case1), 2))
	fmt.Println(reflect.DeepEqual(case1, []int{1, 2, 2}))

	case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(reflect.DeepEqual(removeDuplicates(case2), 5))
	fmt.Println(reflect.DeepEqual(case2, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}))
}
