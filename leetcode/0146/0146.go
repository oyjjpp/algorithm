package leetcode

import (
	"encoding/json"
	"log"
)

// 单个节点
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type DoubleList struct {
	// 链表元素个数
	size int
	// 头尾节点
	head *Node
	tail *Node
}

// InitDoubleList
// 初始化双向链表
func InitDoubleList() *DoubleList {
	list := &DoubleList{
		head: &Node{},
		tail: &Node{},
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	list.size = 0
	return list
}

// AddLast
// 在链表尾部添加节点x，时间复杂度O(1)
func (list *DoubleList) AddLast(x *Node) {
	x.prev = list.tail.prev
	x.next = list.tail

	list.tail.prev = x
	list.tail.prev = x
	list.size++

	// TODO
	log.Println("x节点:", x.key, x.val)
	list.printList()
}

// printList
// 打印当前节点
func (list *DoubleList) printList() {
	log.Println("打印节点", list, list.head.next)
}

// Remove
// 删除链表中x节点（x一定存在）
// 由于是双链表且给得目标Node节点，时间复杂度为O(1)
func (list *DoubleList) Remove(x *Node) {
	// 前一个节点的下一个为x节点得下一个
	x.prev.next = x.next
	// 下一个节点的前一个为x节点的前一个
	x.next.prev = x.prev
	list.size--
}

// RemoveFirst
// 删除链表中第一个节点,并返回该节点，时间复杂度为O(1)
func (list *DoubleList) RemoveFirst() *Node {
	// 校验是否双向链表为空
	if list.head.next == list.tail {
		return nil
	}
	first := list.head.next
	list.Remove(first)
	return first
}

// Size
// 返回链表长度，时间复杂度O(1)
func (list *DoubleList) Size() int {
	return list.size
}

type LRUCache struct {
	hashMap map[int]*Node
	cache   *DoubleList
	cap     int
}

// Constructor
// 初始化
func Constructor(capacity int) LRUCache {
	return LRUCache{
		hashMap: map[int]*Node{},
		cache:   InitDoubleList(),
		cap:     capacity,
	}
}

// makeRecently
// 将key提升为最近使用得
func (this *LRUCache) makeRecently(key int) {
	x, ok := this.hashMap[key]
	// 校验是否存在
	if !ok {
		return
	}
	this.cache.Remove(x)
	this.cache.AddLast(x)
}

// addRecently
// 添加最近使用得元素
func (this *LRUCache) addRecently(key, val int) {
	x := &Node{key: key, val: val}
	// 链表尾部就是最近使用得元素
	this.cache.AddLast(x)
	// 在map中添加key得映射
	this.hashMap[key] = x
}

// deleteKey
// 删除某一个key
func (this *LRUCache) deleteKey(key int) {
	x, ok := this.hashMap[key]
	// 校验是否存在
	if !ok {
		return
	}
	this.cache.Remove(x)
	delete(this.hashMap, key)
}

// removeLeastRecently
// 删除最久未使用的元素
func (this *LRUCache) removeLeastRecently() {
	x := this.cache.RemoveFirst()
	if x == nil {
		return
	}
	key := x.key
	delete(this.hashMap, key)
}

// Get
func (this *LRUCache) Get(key int) int {
	node, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	this.makeRecently(key)
	return node.val
}

// Put
func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.hashMap[key]; ok {
		// 删除旧的数据
		this.deleteKey(key)
		// 新插入的数据为最近使用的数据
		this.addRecently(key, value)
	}

	log.Println("add eme", key, value, this.cap, this.cache.Size())
	if this.cap == this.cache.Size() {
		// 删除最近未使用的元素
		this.removeLeastRecently()
		// TODO
		log.Println("add eme", key, value, this.cap, this.cache.Size())
		test, _ := json.Marshal(this.hashMap)
		log.Println(string(test))
	}
	// 添加为最近使用的元素
	this.addRecently(key, value)
}
