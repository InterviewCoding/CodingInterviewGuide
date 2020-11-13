package main

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/11/2 3:27 下午
 * @Desc: 合并两个有序单链表
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//代码段1
	if l1 == nil {
		return l2
	}

	//代码段2
	if l2 == nil {
		return l1
	}

	dummyNode := &ListNode{}
	cur := dummyNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next, l1, cur = l1, l1.Next, l1
		} else {
			cur.Next, l2, cur = l2, l2.Next, l2
		}
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummyNode.Next
}
