package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。

func PrintListFromTailToHead(head *proto.ListNode) (data []int) {
	if head == nil {
		return
	}
	data = PrintListFromTailToHead(head.Next)
	data = append(data, head.Val)
	return
}