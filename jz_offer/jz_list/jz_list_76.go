package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。 
// 例如，链表 1->2->3->3->4->4->5  处理后为 1->2->5

func deleteDuplication(pHead *proto.ListNode ) *proto.ListNode {
	var res = new(proto.ListNode)
	res.Next = pHead
    cur := res
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == temp {
				cur.Next = cur.Next.Next
			}
        } else {
            cur = cur.Next
        }
	}
    return res.Next
}
