package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//---------------------------------------递归实现二叉树的先序遍历---------------------------------------
//✅代码
func PreorderTraverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	ret = append(ret, root.Val)
	if root.Left != nil {
		ret = append(ret, PreorderTraverse(root.Left)...)
	}

	if root.Right != nil {
		ret = append(ret, PreorderTraverse(root.Right)...)
	}

	return ret
}

//✅
//另外一种书写二叉树先序遍历的方式，不知道是否正确
//验证之后是正确的，因为左子树为空所以返回nil，而nil...将会不执行
func PreorderTraverseII(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	ret = append(ret, root.Val)
	ret = append(ret, PreorderTraverse(root.Left)...)
	ret = append(ret, PreorderTraverse(root.Right)...)

	return ret
}

//❌
//将3个append代码压缩到一句
//压缩到一句是不可以的，
//func PreorderTraverseIII(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	var ret []int
//	ret = append(ret, root.Val, PreorderTraverseIII(root.Left)..., PreorderTraverseIII(root.Right)...)
//
//	return ret
//}

//---------------------------------------递归实现二叉树的中序遍历---------------------------------------
func InorderTraverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	ret = append(ret, InorderTraverse(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, InorderTraverse(root.Right)...)
	return ret
}

//---------------------------------------递归实现二叉树的后序遍历---------------------------------------
func PostorderTraverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	ret = append(ret, PostorderTraverse(root.Left)...)
	ret = append(ret, PostorderTraverse(root.Right)...)
	ret = append(ret, root.Val)

	return ret
}

//---------------------------------------非递归实现二叉树的先序遍历---------------------------------------
//✅
//思路：使用栈来求解，初始化的时候若根不为空，则将根加入到栈中，
//之后，每次遍历到一个节点将值加入到结果中并弹出，然后将右子树加入到栈中，之后将左子树加入到栈中
func PreorderTraverseNoRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		//首先从栈中弹出一个节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//将当前节点的值加入到我们要返回的结果中
		ret = append(ret, node.Val)

		//将当前节点的右节点加入到栈中
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		//将当前节点的左节点加入到栈中
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ret
}

//---------------------------------------非递归实现二叉树的中序遍历---------------------------------------
//✅
func InorderTraverseNoRecursion(root *TreeNode) []int {
	//如果为空，直接返回
	if root == nil {
		return nil
	}

	var ret []int
	var stack []*TreeNode
	cur := root

	for cur != nil || len(stack) != 0 {
		if cur != nil {
			//说明当前节点是第一次遍历，直接加入到栈中
			stack = append(stack, cur)
			//之后移动到该节点的左子树节点
			cur = cur.Left
		} else {
			//首先从栈中弹出一个节点
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//将当前节点的值加入到结果中
			ret = append(ret, cur.Val)
			//之后移动到当前节点的右节点
			cur = cur.Right
		}
	}

	return ret
}

//---------------------------------------非递归实现二叉树的后序遍历---------------------------------------
//✅
//方法一：使用两个栈实现二叉树的非递归后序遍历
func PostorderTraverseNoRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	var stack2 []*TreeNode

	//首先我们将root加入到stack1中
	stack1 := []*TreeNode{root}

	//如果stack1不为空
	for len(stack1) != 0 {
		cur := stack1[len(stack1)-1]
		//将当前节点从栈1移除并且将当前节点加入到栈2
		stack1 = stack1[:len(stack1)-1]

		//将当前节点的左右子节点分别加入到栈1

		//注意点1：这里有可能左右子节点有可能为空
		if cur.Left != nil {
			stack1 = append(stack1, cur.Left)
		}

		if cur.Right != nil {
			stack1 = append(stack1, cur.Right)
		}

		//之后将当前节点加入到栈2中
		stack2 = append(stack2, cur)
	}

	//最后我们从stack2弹出的顺序就是我们后序遍历得到的结果
	for len(stack2) != 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		ret = append(ret, node.Val)
	}

	return ret
}

//方法二：只使用一个栈实现二叉树的非递归后序遍历
//✅
//思路：使用一个栈和两个变量，h代表上次访问并删除的节点，c代表当前节点
func PostorderTraverseNoRecursionII(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	stack := []*TreeNode{root}

	//初始化的时候h置为root，将c置为nil
	h := root
	var c *TreeNode

	//如果栈不为空
	for len(stack) != 0 {
		c = stack[len(stack)-1]
		//如果当前节点的左子树不为空，并且左右子树都不等于h，说明左子树没有遍历过，将左节点加入栈中
		if c.Left != nil && h != c.Left && h != c.Right {		//也就是第一次遍历该节点的时候
			stack = append(stack, c.Left)
		} else if c.Right != nil && c.Right != h { //如果当前节点的右子树不为空，且不等于h，说明没有遍历过，则将右节点加入到栈中，也就是第二次遍历该节点的时候
			stack = append(stack, c.Right)
		} else { //否则，弹出节点，并加入到结果中，也就是第三次遍历该节点的时候
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, node.Val)
			//并且将上一次访问并且打印过的节点重置为node
			h = node
		}
	}

	return ret
}

func main() {

	right := &TreeNode{
		Val: 4,
	}
	//建立一个树
	root := &TreeNode{
		Val: 3,
	}

	root.Right = right

	fmt.Println(PreorderTraverse(root))
	fmt.Println(PreorderTraverseII(root))
	fmt.Println("123")
}
