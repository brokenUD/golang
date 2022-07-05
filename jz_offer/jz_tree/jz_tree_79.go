package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	"math"
	// "math"
)

// 描述
// 输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
// 在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树
// 平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。

// 样例二叉树如图，为一颗平衡二叉树
// 注：我们约定空树是平衡二叉树。
// 数据范围：n \le 100n≤100,树上节点的val值满足 0 \le n \le 10000≤n≤1000
// 要求：空间复杂度O(1)O(1)，时间复杂度 O(n)O(n)

func IsBalanced_Solution(pRoot *proto.TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return int(math.Abs(float64(depth(pRoot.Left))-float64(depth(pRoot.Right)))) <= 1 && IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
}

func depth(pRoot *proto.TreeNode) int {
	if pRoot == nil {
		return 0
	}
	return int(math.Max(float64(depth(pRoot.Left)), float64(depth(pRoot.Right)))) + 1
}
