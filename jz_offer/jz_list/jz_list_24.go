package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 给定一个单链表的头结点pHead(该头节点是有值的，比如在下图，它的val是1)，长度为n，反转该链表后，返回新链表的表头。

func ReverseList( pHead *proto.ListNode ) *proto.ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	var ans *proto.ListNode
	for pHead != nil {
		temp := pHead.Next
		pHead.Next = ans
		ans = pHead
		pHead = temp
	} 
	return ans
}