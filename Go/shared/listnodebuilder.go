package shared

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(items []int) *ListNode {
	var ln *ListNode
	items = reverse(items)
	for i := range items {
		ln = &ListNode{i, ln}
	}
	return ln
}

// https://dev.to/dsysd_dev/reverse-a-generic-list-in-golang-2fd4
func reverse[T any](list []T) []T {
	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}
