package main

/**
 * Definition for singly-linked list.
 */

/*
type ListNode struct {
	Val  int
	Next *ListNode
}
*/

// See https://leetcode.cn/problems/merge-in-between-linked-lists/description/ for more details
/*
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var (
		head = list1
		idx  = 0
		pre  *ListNode
		next *ListNode
	)

	for list1 != nil {
		if idx == a-1 {
			pre = list1
		}
		if idx == b+1 {
			next = list1
			break
		}
		idx++
		list1 = list1.Next
	}

	pre.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = next

	return head
}
*/

/*
	简单模拟即可
*/
