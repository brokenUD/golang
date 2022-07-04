package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）

func FindFirstCommonNode( pHead1 *proto.ListNode ,  pHead2 *proto.ListNode ) *proto.ListNode {
	l1, l2 := pHead1, pHead2
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = pHead2
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = pHead1
		}
	}
	return l1
}