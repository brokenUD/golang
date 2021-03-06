package jz_tree

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。

// 提示:
// 1.vin.length == pre.length
// 2.pre 和 vin 均无重复元素
// 3.vin出现的元素均出现在 pre里
// 4.只需要返回根结点，系统会自动输出整颗树做答案对比
// 数据范围：n \le 2000n≤2000，节点的值 -10000 \le val \le 10000−10000≤val≤10000
// 要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

func reConstructBinaryTree(pre []int, vin []int) *proto.TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &proto.TreeNode{
		Val: pre[0],
	}
	var tmp int
	for tmp = 0; tmp < len(vin); tmp++ {
		if vin[tmp] == pre[0] {
			break
		}
	}
	root.Left = reConstructBinaryTree(pre[1:tmp+1], vin[:tmp])
	root.Right = reConstructBinaryTree(pre[tmp+1:], vin[tmp+1:])
	return root
}
