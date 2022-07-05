package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定一个二叉树其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。
// 注意，树中的结点不仅包含左右子结点，同时包含指向父结点的next指针。
// 下图为一棵有9个节点的二叉树。树中从父节点指向子节点的指针用实线表示，从子节点指向父节点的用虚线表示

func GetNext(pNode *proto.TreeLinkNode) *proto.TreeLinkNode {
	if pNode == nil {
		return nil
	}
	if pNode.Right != nil {
		right := pNode.Right
		for right.Left != nil {
			right = right.Left
		}
		return right
	}
	if pNode.Next != nil && pNode.Next.Left == pNode {
		return pNode.Next
	}
	// 牛逼
	if pNode.Next != nil {
		next := pNode.Next
		for next.Next != nil && next.Next.Right == next {
			next = next.Next
		}
		return next.Next
	}
	return nil
}
