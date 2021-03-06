package main

import (
	"fmt"
	"strconv"
	"strings"
)

//✅
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//---------------------------------------二叉树的序列化和反序列化---------------------------------------


//-------------------------------------------------------第1种写法-------------------------------------------------------

//✅
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

	//注意点：由于切片在执行过程中有可能会因为增加和删除元素而造成切片不是原来那个切片，但是我们递归回去的时候还是指向原来的切片，因此会有问题
	//所以这里我们传递的是切片的地址
	//因为切片扩容可能会生成一个新的底层数组，并且由于切片移除了元素，因此对应的头部地址一定会改变，所以会造成地址的改变
	root := ReconstructTreeFromPreorder(&strs)
	return root
}

//根据我们分割后的字符串建立二叉树
func ReconstructTreeFromPreorder(strs *[]string) *TreeNode {
	if (*strs)[0] == "#" {
		(*strs) = (*strs)[1:]
		return nil
	}

	//首先将该值对应的字符串转换为int
	val, _ := strconv.Atoi((*strs)[0])
	//建立一个针对于该值的节点
	node := &TreeNode{
		Val: val,
	}

	//去掉我们建立过的节点的值
	(*strs) = (*strs)[1:]
	//之后进行递归建立左右子树
	node.Left = ReconstructTreeFromPreorder(strs)
	node.Right = ReconstructTreeFromPreorder(strs)

	return node
}

//-------------------------------------------------------第2种写法-------------------------------------------------------

//✅
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


//根据层次遍历重建二叉树
func LevelOrderReconstruct(strs []string) *TreeNode {
	var head *TreeNode
	//一进来首先建立根节点
	if strs[0] == "#" {
		return nil
	} else {
		val, _ := strconv.Atoi(strs[0])
		head = &TreeNode{
			Val: val,
		}
	}

	//用于表示我们当前遍历到的节点是层次遍历结果中的第几个节点
	//因为头节点已经建立完，说明下次我们需要从strs中的index为1的位置开始建立节点
	index := 1

	//使用一个队列将我们后面需要建立子树的非空节点加入进来，
	queue := []*TreeNode{head}
	for len(queue) != 0 {
		//说明当前开始建立cur的左右子树
		cur := queue[0]
		queue = queue[1:]

		//建立左子树
		if strs[index] == "#" {
			cur.Left = nil
		} else {
			val, _ := strconv.Atoi(strs[index])
			cur.Left = &TreeNode{
				Val: val,
			}
			queue = append(queue, cur.Left)
		}
		index++

		//建立右子树
		if strs[index] == "#" {
			cur.Right = nil
		} else {
			val, _ := strconv.Atoi(strs[index])
			cur.Right = &TreeNode{
				Val: val,
			}
			queue = append(queue, cur.Right)
		}
		index++
	}

	return head
}

//-------------------------------------------------------第3种写法（改进）-------------------------------------------------------
//✅
//对方法二的改进
//将臃肿的代码进行函数的重构，重构成一个简单的函数：generateNodeFromString
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

//根据层次遍历重建二叉树
func LevelOrderReconstruct(strs []string) *TreeNode {
	//一进来首先建立根节点
	head := generateNodeFromString(strs[0])

	//这里要注意判断一下head是否为空
	//注意点：判断head是否为空，如果为空就推出，否则后面会报错的
	if head == nil {
		return nil
	}

	//用于表示我们当前遍历到的节点是层次遍历结果中的第几个节点
	//因为头节点已经建立完，说明下次我们需要从strs中的index为1的位置开始建立节点
	index := 1

	//使用一个队列将我们后面需要建立子树的非空节点加入进来，
	queue := []*TreeNode{head}
	for len(queue) != 0 {
		//说明当前开始建立cur的左右子树
		cur := queue[0]
		queue = queue[1:]

		//建立左子树
		cur.Left = generateNodeFromString(strs[index])
		index++

		//建立右子树
		cur.Right = generateNodeFromString(strs[index])
		index++

		//如果左子树不为空，加入到队列中
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		//如果右子树不为空，加入到队列中
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return head
}

func generateNodeFromString(val string) *TreeNode {
	if val == "#" {
		return nil
	}

	temp, _ := strconv.Atoi(val)
	return &TreeNode{
		Val: temp,
	}
}

func main() {
	//建立一颗二叉树
	left := &TreeNode{
		Val: 2,
	}

	leftR := &TreeNode{
		Val: 4,
	}

	rightL := &TreeNode{
		Val: 9,
	}

	rightR := &TreeNode{
		Val: 23,
	}

	right := &TreeNode{
		Val: 3,
	}

	left.Right = leftR
	right.Left = rightL
	right.Right = rightR

	root := &TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}
	///1!2!#!4!#!#!3!9!#!#!23!#!#!
	fmt.Println(PreorderSerialize(root))

	r := PreorderDeserialize(PreorderSerialize(root))
	fmt.Println(r.Val)
	fmt.Println(r.Left.Val)
	fmt.Println(r.Left.Right.Val)
	fmt.Println(r.Right.Val)
	fmt.Println(r.Right.Left.Val)
	fmt.Println(r.Right.Right.Val)
}
