package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定一棵结点数为n 二叉搜索树，请找出其中的第 k 小的TreeNode结点值。
// 1.返回第k小的节点值即可
// 2.不能查找的情况，如二叉树为空，则返回-1，或者k大于n等等，也返回-1
// 3.保证n个节点的值不一样

// 数据范围： 0 \le n \le10000≤n≤1000，0 \le k \le10000≤k≤1000，树上每个结点的值满足0 \le val \le 10000≤val≤1000
// 进阶：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

func KthNode( proot *proto.TreeNode ,  k int ) int {
	if proot == nil {
		return -1
	}
	var count, ans int
	ans, _ = midOrder(proot, k, count)
    return ans
}

func midOrder(root *proto.TreeNode, k, count int)(int, int) {
	if root == nil ||count > k {
		return -1, count
	}
    var tmp int
	tmp, count = midOrder(root.Left, k, count)
	if tmp > 0 {
		return tmp, count
	}
	count++
	if count == k {
		return root.Val, count
	}
	tmp, count = midOrder(root.Right, k ,count)
	if tmp > 0 {
		return tmp, count
	}
	return -1, count
}