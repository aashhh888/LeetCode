package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(reflect.DeepEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}))
	fmt.Println(reflect.DeepEqual(twoSum([]int{3, 2, 4}, 6), []int{1, 2}))
	fmt.Println(reflect.DeepEqual(twoSum([]int{3, 3}, 6), []int{0, 1}))
}
