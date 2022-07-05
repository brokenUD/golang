package jz_list

import (
	// "brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回 true ,否则返回 false 。假设输入的数组的任意两个数字都互不相同。
// 数据范围： 节点数量 0 \le n \le 10000≤n≤1000 ，节点上的值满足 1 \le val \le 10^{5}1≤val≤10^5，保证节点上的值各不相同
// 要求：空间复杂度 O(n)O(n) ，时间时间复杂度 O(n^2)O(n^2)
// 提示：
// 1.二叉搜索树是指父亲节点大于左子树中的全部节点，但是小于右子树中的全部节点的树。
// 2.该题我们约定空树不是二叉搜索树
// 3.后序遍历是指按照 “左子树-右子树-根节点” 的顺序遍历

func VerifySquenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		return false
	}
	n := len(sequence)
	return check(sequence, 0, n-1)
}

func check(sequence []int, l int, r int) bool {
	if l >= r {
		return true
	}
	root := sequence[r]
	var j int
	for j = r; j > l; j-- {
		if sequence[j] < root {
			break
		}
	}
	// <没有=
	for i := l; i < j; i++ {
		if sequence[i] > root {
			return false
		}
	}
	return check(sequence, l, j) && check(sequence, j+1, r-1)
}
