package jz_list

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 输入两棵二叉树A，B，判断B是不是A的子结构。（我们约定空树不是任意一个树的子结构）
// 假如给定A为{8,8,7,9,2,#,#,#,#,4,7}，B为{8,9,2}，2个树的结构如下，可以看出B是A的子结构

// 提示:
// 数据范围:
// 0 <= A的节点个数 <= 10000
// 0 <= B的节点个数 <= 10000

func HasSubtree(pRoot1 *proto.TreeNode, pRoot2 *proto.TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	return isSubtree(pRoot1, pRoot2) || HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func isSubtree(pRoot1 *proto.TreeNode, pRoot2 *proto.TreeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return isSubtree(pRoot1.Left, pRoot2.Left) && isSubtree(pRoot1.Right, pRoot2.Right)
}
