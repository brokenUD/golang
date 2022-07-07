package jz_tree

import (
	"brokenUd/golang/jz_offer/proto"
	"math"
)

// 描述
// 输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，
// 最长路径的长度为树的深度，根节点的深度视为 1 。

func TreeDepth( pRoot *proto.TreeNode ) int {
	if pRoot == nil {
		return 0
	}	
	l := TreeDepth(pRoot.Left)
	r := TreeDepth(pRoot.Right)
	return int(math.Max(float64(l), float64(r))) + 1
}