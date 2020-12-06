package main

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func getMaxBST(root *TreeNode) *TreeNode {
	return process(root).Head
}

type RetType struct {
	//该节点下面的最大搜索二叉子树的头结点
	Head     *TreeNode
	Count    int //该子树的节点个数
	Max, Min int //该子树对应的最大值和最小值
}

//❌
//自己写的错误代码：如果输入一棵树，右子树不为空，但是是一颗二叉搜索树，只会返回一个右子树的根节点而不是整棵树的根节点
//递归函数
func process(root *TreeNode) RetType {
	if root == nil {
		return RetType{
			Head:  nil,
			Count: 0,
			Max:   math.MinInt32,
			Min:   math.MaxInt32,
		}
	}

	left, right := process(root.Left), process(root.Right)

	//如果左子树不是直接返回右子树
	if left.Head == nil {
		return right
	}

	//如果右子树不是直接返回左子树
	if right.Head == nil {
		return left
	}

	//说明左右子树不为空，并且里面一定有二叉搜索子树，我们只需要判断根节点是否可以合并进去
	if left.Head == root.Left && right.Head == root.Right && root.Val > left.Max && root.Val < right.Min {
		return RetType{
			Head:  root,
			Count: 1 + left.Count + right.Count,
			Max:   right.Max,
			Min:   left.Min,
		}
	}

	//说明不是，需要判断左右子树哪个数量多返回哪个
	if left.Count > right.Count {
		return left
	} else {
		return right
	}
}

//自己重新按照自己的理解写一个函数
func process_mine(root *TreeNode) RetType {
	//如果当前节点为空
	if root == nil {
		return RetType{
			Head: nil,
			Count: 0,
			Max: math.MaxInt32,
			Min: math.MinInt32,
		}
	}

	var head *TreeNode
	var count, curMax, curMin int


}


func processII(root *TreeNode) RetType {
	if root == nil {
		return RetType{
			Head:  nil,
			Count: 0,
			Max:   math.MinInt32,
			Min:   math.MaxInt32,
		}
	}

	//得到两个子树的全部信息
	left, right := processII(root.Left), processII(root.Right)

	//更新当前节点的信息
	curMin := min(root.Val, min(left.Min, right.Min))
	curMax := max(root.Val, max(left.Max, right.Max))
	count := max(left.Count, right.Count)
	var head *TreeNode
	if left.Count > right.Count {
		head = left.Head
	} else {
		head = right.Head
	}

	//判断是否满足第3种可能，即最大二叉搜索子树可以将根合并进去
	if left.Head == root.Left && right.Head == root.Right && root.Val > left.Max && root.Val < right.Min {
		count = left.Count + right.Count + 1
		head = root
	}

	return RetType{
		Head:  head,
		Count: count,
		Max:   curMax,
		Min:   curMin,
	}


}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
