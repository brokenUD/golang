package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 操作给定的二叉树，将其变换为源二叉树的镜像。
// 数据范围：二叉树的节点数 0 \le n \le 10000≤n≤1000 ， 二叉树每个节点的值 0\le val \le 10000≤val≤1000
// 要求： 空间复杂度 O(n)O(n) 。本题也有原地操作，即空间复杂度 O(1)O(1) 的解法，时间复杂度 O(n)O(n)

func Mirror(pRoot *proto.TreeNode) *proto.TreeNode {
	if pRoot == nil {
		return nil
	}
	tmp := pRoot.Left
	pRoot.Left = pRoot.Right
	pRoot.Right = tmp
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}
