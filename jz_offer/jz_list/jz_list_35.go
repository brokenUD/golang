package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
)

// 描述
// 输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针random指向一个随机节点），
// 请对此链表进行深拷贝，并返回拷贝后的头结点。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）。 
// 下图是一个含有5个结点的复杂链表。图中实线箭头表示next指针，虚线箭头表示random指针。为简单起见，指向null的指针没有画出。

func Clone( head *proto.RandomListNode ) *proto.RandomListNode {
	if head == nil {
		return nil 
	}
	cur := head
	for cur != nil {
		temp := &proto.RandomListNode{
			Label: cur.Label,
			Next: cur.Next,
		}
		cur.Next = temp
		cur = temp.Next
	}
	old, clone, ret := head, head.Next, head.Next
	for old != nil {
		if old.Random != nil {
			clone.Random = old.Random.Next
		}
		if old.Next != nil {
			old = old.Next.Next
		}
		if clone.Next != nil {
			clone = clone.Next.Next
		}
	}
	old, clone = head, head.Next
	for old != nil {
		if old.Next != nil {
			old.Next = old.Next.Next
		}
		if clone.Next != nil {
			clone.Next = clone.Next.Next
		}
		old = old.Next
		clone = clone.Next
	}
    return ret
}
