package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。

func EntryNodeOfLoop(pHead *proto.ListNode) *proto.ListNode {
	if pHead == nil {
		return pHead
	}
	pLow, pFast := pHead, pHead
	for pFast != nil && pFast.Next != nil {
		pFast = pFast.Next.Next
		pLow = pLow.Next
		if pLow == pFast {
			break
		}
	}
	if pFast == nil || pFast.Next == nil {
		return nil
	}
	pFast = pHead
	for pFast != pLow {
		pFast = pFast.Next
		pLow = pLow.Next
	}
	return pFast
}