package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。返回删除后的链表的头节点。
// 1.此题对比原题有改动
// 2.题目保证链表中节点的值互不相同
// 3.该题只会输出返回的链表和结果做对比，所以若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点

func deleteNode( head *proto.ListNode ,  val int ) *proto.ListNode {
	var res = new(proto.ListNode)
	res.Next = head
	cur := res
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
    return res.Next
}
