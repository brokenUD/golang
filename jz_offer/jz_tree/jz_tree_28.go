package jz_tree

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）

func isSymmetrical(pRoot *proto.TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return compare(pRoot.Left, pRoot.Right)
}

func compare(t1 *proto.TreeNode, t2 *proto.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}
	return compare(t1.Left, t2.Right) && compare(t1.Right, t2.Left)
}
