package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定一个二叉树root和一个整数值 sum ，求该树有多少路径的的节点值之和等于 sum 。
// 1.该题路径定义不需要从根节点开始，也不需要在叶子节点结束，但是一定是从父亲节点往下到孩子节点
// 2.总节点数目为n
// 3.保证最后返回的路径个数在整形范围内(即路径个数小于231-1)

// 数据范围:
// 0<=n<=10000<=n<=1000-10^9<=节点值<=10^9−10 9 <=节点值<=10 9

func FindPath2(root *proto.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return FindPathDfs2(root, sum) + FindPath2(root.Left, sum) + FindPath2(root.Right, sum)
}

func FindPathDfs2(root *proto.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum -= root.Val
	var n int
	if sum == 0 {
		n = 1
	}
	return FindPathDfs2(root.Left, sum) + FindPathDfs2(root.Right, sum) + n
}
