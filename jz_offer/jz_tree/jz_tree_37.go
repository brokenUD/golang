package jz_tree

import (
	"brokenUd/golang/jz_offer/proto"
	"strconv"
	"strings"
	// "math"
)

// 描述
// 请实现两个函数，分别用来序列化和反序列化二叉树，不对序列化之后的字符串进行约束，
// 但要求能够根据序列化之后的字符串重新构造出一棵与原二叉树相同的树。
// 二叉树的序列化(Serialize)是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，
// 从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树等遍历方式来进行修改，
// 序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#）
// 二叉树的反序列化(Deserialize)是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。

func Serialize(root *proto.TreeNode) string {
	if root == nil {
		return "#"
	}
	return string(root.Val) + "," + Serialize(root.Left) + "," + Serialize(root.Right)
}

func Deserialize(s string) *proto.TreeNode {
	str := strings.Split(s, ",")
	return buildTree(&str)
}

func buildTree(s *[]string) *proto.TreeNode {
	if s == nil {
		return nil
	}
	valS := (*s)[0]
	if valS == "#" {
		return nil
	}
	val, _ := strconv.Atoi(valS)
	root := &proto.TreeNode{
		Val: val,
	}
	*s = (*s)[1:]
	root.Left = buildTree(s)
	root.Right = buildTree(s)
	return root
}
