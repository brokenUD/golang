package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
// 1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
// 2.叶子节点是指没有子节点的节点
// 3.路径只能从父节点到子节点，不能从子节点到父节点
// 4.总节点数目为n

// 数据范围：
// 1.树上的节点数满足 0 \le n \le 100000≤n≤10000
// 2.每 个节点的值都满足 |val| \le 1000∣val∣≤1000
// 要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
// 进阶：空间复杂度 O(树的高度)O(树的高度)，时间复杂度 O(n)O(n)

func hasPathSum(root *proto.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumDfs(root, sum)
}

func hasPathSumDfs(root *proto.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSumDfs(root.Left, sum-root.Val) || hasPathSumDfs(root.Right, sum-root.Val)
}