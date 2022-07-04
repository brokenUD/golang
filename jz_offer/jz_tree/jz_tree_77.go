package jz_list

import (
	"brokenUd/golang/jz_offer/proto"

	"git.ixiaochuan.cn/xc_server/xclib/common/utils/queue"
	// "math"
)

// 描述
// 给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）

func Print( pRoot *proto.TreeNode ) [][]int {
	if pRoot == nil {
		return nil
	}
	var ans [][]int
	queue := []*proto.TreeNode{pRoot}
	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int, n)
		for i:=0;i<n;i++{
			if queue[i]==nil{
				continue
			}
			if len(ans)%2==0{
				tmp[i]=queue[i].val
			}
		}
	}
    return
}

/*

//BFS+队列
func Print( root *TreeNode ) [][]int {
    // write code here
    res := [][]int{}
    if root == nil {return nil}
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        tmp := make([]int, n)
        for i := 0; i < n; i++ {
            if queue[i] == nil {
                continue
            }
            if len(res)%2 == 0 {
                tmp[i] = queue[i].Val
            }else {
                tmp[n-1-i] = queue[i].Val
            }
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil  {
                queue = append(queue, queue[i].Right)
            }
        }
        queue = queue[n:]
        res = append(res, tmp)
    }
    return res
}
