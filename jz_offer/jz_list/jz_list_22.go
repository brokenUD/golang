package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
// 如果该链表长度小于k，请返回一个长度为 0 的链表。

func FindKthToTail( pHead *proto.ListNode ,  k int ) *proto.ListNode {
	var ans *proto.ListNode
	if pHead == nil {
		return ans
	}
	pLow, pFast := pHead, pHead
	for i := 0; i < k; i++ {
		if pFast == nil {
			return ans
		}
		pFast = pFast.Next
	}
	for pFast != nil {
		pFast = pFast.Next
		pLow = pLow.Next
	}
    return pLow
}