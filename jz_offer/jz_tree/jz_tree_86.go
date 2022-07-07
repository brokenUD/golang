package jz_tree

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
// 数据范围：树上节点数满足 1 \le n \le 10^5 \1≤n≤10 5, 节点值val满足区间 [0,n)
// 要求：时间复杂度 O(n)O(n)
// 注：本题保证二叉树中每个节点的val值均不相同。

func lowestCommonAncestor(root *proto.TreeNode, o1 int, o2 int) int {
	ans := helper(root, o1, o2)
	return ans.Val
}

func helper(root *proto.TreeNode, o1, o2 int) *proto.TreeNode {
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	left := helper(root.Left, o1, o2)
	right := helper(root.Right, o1, o2)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
