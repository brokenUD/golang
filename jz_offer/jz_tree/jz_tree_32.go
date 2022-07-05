package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 不分行从上往下打印出二叉树的每个节点，同层节点从左至右打印。例如输入{8,6,10,#,#,2,1}，
// 如以下图中的示例二叉树，则依次打印8,6,10,2,1(空节点不打印，跳过)，请你将打印的结果存放到一个数组里面，返回

func PrintFromTopToBottom(root *proto.TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue = make([]*proto.TreeNode, 0)
	var ans = make([]int, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			ans = append(ans, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	return ans
}
