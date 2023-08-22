package linkedlist

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false 。*/

// https://leetcode.cn/problems/linked-list-cycle/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	dummyNode := head
	for dummyNode != nil {
		if dummyNode.Val != math.MaxInt {
			dummyNode.Val = math.MaxInt
			dummyNode = dummyNode.Next
			continue
		}
		return true
	}
	return false
}
