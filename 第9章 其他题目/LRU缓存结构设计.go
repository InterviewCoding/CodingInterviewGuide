package main

/**
 * @Author: yirufeng
 * @Date: 2020/11/21 11:33 上午
 * @Desc: 设计LRU缓存结构
思路：map + 双向链表
 **/

//✅
type DoubleLinkedListNode struct {
	Key, Val int
	Next, Prev *DoubleLinkedListNode
}

//lruCache的结构
type LRUCache struct {
	Cap int
	Size int
	Head, Tail *DoubleLinkedListNode
	CacheMap map[int]*DoubleLinkedListNode
}

//构造函数
func Constructor(capacity int) LRUCache {

	//建立虚拟的头和尾节点
	head, tail := &DoubleLinkedListNode{}, &DoubleLinkedListNode{}
	head.Prev, head.Next, tail.Prev, tail.Next = tail, tail, head, head
	//建立map
	tempMap := make(map[int]*DoubleLinkedListNode)
	return LRUCache{
		Cap: capacity,
		Size: 0,
		Head: head,
		Tail: tail,
		CacheMap: tempMap,
	}
}

//移除节点
func (this *LRUCache) removeNode(node *DoubleLinkedListNode) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

//加入到头部
func (this *LRUCache) addToHead(node *DoubleLinkedListNode) {
	this.Head.Next, this.Head.Next.Prev, node.Prev, node.Next = node, node, this.Head, this.Head.Next
}

//获取值
func (this *LRUCache) Get(key int) int {
	//找到了该节点
	if node, ok := this.CacheMap[key]; ok {
		//移除该节点
		this.removeNode(node)
		//提升为最近使用
		this.addToHead(node)
		//返回该节点的值
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//找到了该节点
	if node, ok := this.CacheMap[key]; ok {
		//修改值
		node.Val = value
		//移除该节点
		this.removeNode(node)
		//提升为最近使用
		this.addToHead(node)
	} else {
		//创建节点
		node := &DoubleLinkedListNode{
			Key: key,
			Val: value,
		}
		//注意点1：记得让大小加上1
		this.Size++
		//提升为最近使用
		this.addToHead(node)
		//注意点2：将节点加入到map中
		this.CacheMap[key] = node
		//如果大小超过了容量就移除最近很久未使用的
		if this.Size > this.Cap {
			//这里有因为从链表中移除了，如果再从map中获取尾巴前面一个节点就不对了，因为获取的不是我们要删除的节点，而是删除之后尾巴前面的节点
			//❌
			/*//从链表中移除
			this.removeNode(this.Tail.Prev)
			//注意点3：从map中移除
			delete(this.CacheMap, this.Tail.Prev.Key)
			//注意点4：删除节点记得大小减去1
			this.Size--*/

			//✅
			node := this.Tail.Prev
			this.removeNode(node)
			delete(this.CacheMap, node.Key)
			this.Size--

		}
	}
}
