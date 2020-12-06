package main

/**
 * @Author: yirufeng
 * @Date: 2020/11/18 10:26 下午
 * @Desc: 二叉树的按层打印与ZigZag打印
 **/
//-------------------------------------------------------按层打印-------------------------------------------------------
//-------------------------------------------------------第1种写法-------------------------------------------------------
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		var curLevelRet []int
		for length != 0 {
			//弹出一个节点
			node := queue[0]
			queue = queue[1:]

			//将结果加入到当前层中
			curLevelRet = append(curLevelRet, node.Val)

			//加入子树
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ret = append(ret, curLevelRet)
	}
	return ret
}

//-------------------------------------------------------第2种写法-------------------------------------------------------
//使用两个变量分别表示当前行还剩多少个节点没有打印以及下一行要打印的节点个数
func levelOrderII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var ret [][]int

	//分别表示当前层还剩多少个节点以及下一层有多少个节点需要打印
	count := 1
	next := 0

	for len(queue) != 0 {
		var curLevlelRet []int
		//如果当前层要打印的节点数不等于0
		for count != 0 {
			node := queue[0]
			queue = queue[1:]
			//注意减去1
			count -=1
			if node.Left != nil {
				queue = append(queue, node.Left)
				next++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				next++
			}
			curLevlelRet = append(curLevlelRet, node.Val)
		}
		//将当前层的结果加入到里面
		ret = append(ret, curLevlelRet)
		//将当前层要打印节点的个数以及下一层需要打印的节点个数重置
		count, next = next, 0
	}
	return ret
}

//-------------------------------------------------------第3种写法-------------------------------------------------------
//参考左神书上写法
//使用两个变量分别表示当前行最后一个节点以及下一行的最后一个节点
func levelOrderIII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var ret [][]int

	//使用两个变量分别表示当前行最后一个节点以及下一行的最后一个节点
	var last, nextLast *TreeNode
	last = root
	var curLevelRet []int

	for len(queue) != 0 {
		//每次从队列中弹出一个节点
		cur := queue[0]
		queue = queue[1:]
		curLevelRet = append(curLevelRet, cur.Val)

		//首先将不为空的左右子树加入到队列中
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			nextLast = cur.Left
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
			nextLast = cur.Right
		}

		//说明要换行了
		if cur == last {
			ret = append(ret, curLevelRet)
			curLevelRet = []int{}
			if len(queue) != 0 {  //说明下一行还有
				last = nextLast
			}
		}
	}

	return ret
}

//-------------------------------------------------------zigzag打印第一种写法---------------------------------------------
//思路：使用双栈进行控制
func zigzagOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	s1, s2 := []*TreeNode{}, []*TreeNode{}

	return
}

func main() {

}
