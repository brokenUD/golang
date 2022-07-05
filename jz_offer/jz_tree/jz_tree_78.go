package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定一个节点数为 n 二叉树，要求从上到下按层打印二叉树的 val 值，
// 同一层结点从左至右输出，每一层输出一行，将输出的结果存放到一个二维数组中返回。

func Print2(pRoot *proto.TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	var ans [][]int
	var queue = make([]*proto.TreeNode, 0)
	queue = append(queue, pRoot)
	for len(queue) > 0 {
		n := len(queue)
		var tmp []int
		for i := 0; i < n; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ans = append(ans, tmp)
		queue = queue[n:]
	}
	return ans
}
