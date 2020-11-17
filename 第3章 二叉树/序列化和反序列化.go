package main

import (
	"strconv"
	"strings"
)

type TreeNdoe struct {
	Val         int
	Left, Right *TreeNode
}

//---------------------------------------二叉树的序列化和反序列化---------------------------------------
//方法一：使用先序遍历进行序列化和反序列化
//进行序列化
func PreorderSerialize(root *TreeNode) string {
	if root == nil {
		return "#!"
	}

	var ret string

	ret += strconv.Itoa(root.Val) + "!"
	ret += PreorderSerialize(root.Left)
	ret += PreorderSerialize(root.Right)
	return ret
}

//进行反序列化
func PreorderDeserialize(ret string) *TreeNode {
	strs := strings.Split(ret, "!")

	root := ReconstructTreeFromPreorder(strs)

	return root

}

//根据我们分割后的字符串建立二叉树
func ReconstructTreeFromPreorder(strs []string) *TreeNode {
	if strs[0] == "#" {
		return nil
	}

	//首先将该值对应的字符串转换为int
	val, _ := strconv.Atoi(strs[0])
	//建立一个针对于该值的节点
	node := &TreeNode{
		Val: val,
	}

	//去掉我们建立过的节点的值
	strs = strs[1:]
	//之后进行递归建立左右子树
	node.Left = ReconstructTreeFromPreorder(strs)
	node.Right = ReconstructTreeFromPreorder(strs)

	return node
}

//方法二：使用层次遍历进行序列化和反序列化
func LevelOrderSerialize(root *TreeNode) string {
	//首先如果根节点为空，我们直接返回一个空字符串即可
	if root == nil {
		return "#!"
	}

	var ret string
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		//首先弹出一个节点
		node := queue[0]
		queue = queue[1:]

		//如果弹出的节点不为空，我们直接将值转换为字符串并加入其中
		if node != nil {
			ret += strconv.Itoa(node.Val) + "!"
		} else {
			ret += "#!"
			//这里因为是当前节点为空，因此加入到#!到字符串之后我们需要重新循环，不可以执行下面的代码
			continue
		}
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	return ret
}

func LevelOrderDeserialize(str string) *TreeNode {

	strs := strings.Split(str, "!")
	root := LevelOrderReconstruct(strs)
	return root
}

func LevelOrderReconstruct(strs []string) *TreeNode {

	//说明当前节点没有值
	if strs[0] == "#" {
		return nil
	}

	val, _ := strconv.Atoi(strs[0])
	node := &TreeNode{
		Val: val,
	}
	strs = strs[1:]
	node.Left = LevelOrderReconstruct(strs)
	node.Right = LevelOrderReconstruct(strs)

	return node
}
