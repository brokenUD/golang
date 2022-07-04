package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
// 数据范围： 0 \le n \le 10000≤n≤1000，-1000 \le 节点值 \le 1000−1000≤节点值≤1000
// 要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

func Merge( pHead1 *proto.ListNode ,  pHead2 *proto.ListNode ) *proto.ListNode {
	var head = new(proto.ListNode)
	cur := head
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			head.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			head.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}
    if pHead1 == nil {
		head.Next = pHead2
	}
	if pHead2 == nil {
		head.Next = pHead1
	}
	return head.Next
}
