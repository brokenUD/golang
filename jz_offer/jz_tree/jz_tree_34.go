package jz_tree

import (
	"brokenUd/golang/jz_offer/proto"
	// "math"
)

// 描述
// 输入一颗二叉树的根节点root和一个整数expectNumber，找出二叉树中结点值的和为expectNumber的所有路径。
// 1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
// 2.叶子节点是指没有子节点的节点
// 3.路径只能从父节点到子节点，不能从子节点到父节点
// 4.总节点数目为n

// 数据范围:
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= 节点值 <= 1000
// -1000 <= expectNumber <= 1000

func FindPath(root *proto.TreeNode, expectNumber int) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var tmp []int
	ans = FindPathDfs(root, expectNumber, tmp, ans)
	return ans
}

func FindPathDfs(root *proto.TreeNode, expectNumber int, tmp []int, ans [][]int) [][]int {
	if root == nil {
		return ans
	}
	expectNumber -= root.Val
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil && expectNumber == 0 {
		ans = append(ans, tmp)
	}
	ans1 := FindPathDfs(root.Left, expectNumber, tmp, ans)
	ans2 := FindPathDfs(root.Right, expectNumber, tmp, ans)
	ans = append(ans, ans1...)
	ans = append(ans, ans2...)
	return ans
}
