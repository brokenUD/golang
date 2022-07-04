package proto

type ListNode struct {
	Val  int
	Next *ListNode
}

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
