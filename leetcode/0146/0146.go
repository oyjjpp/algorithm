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
	// 链接头尾节点
	list.head.next = list.tail
	list.tail.prev = list.head
	list.size = 0
	return list
}

// AddLast
// 在链表尾部添加节点x，时间复杂度O(1)
// 增加一个新的节点
// 1、首先要确认新增加节点的，前后节点
// 2、更改前后节点，摘除原有的关联之后，与x节点建立联系
func (list *DoubleList) AddLast(x *Node) {
	// 先将尾节点的前一个元素赋值给x的前节点，x后节点赋值为尾节点
	x.prev = list.tail.prev
	x.next = list.tail

	// 尾节点的前一个节点的下一个节点更改尾x
	list.tail.prev.next = x
	list.tail.prev = x

	// 增加元素个数
	list.size++
}

// Remove
// 删除链表中x节点（x一定存在）
// 由于是双链表且给得目标Node节点，时间复杂度为O(1)
// 删除一个节点主要是将删除节点的前一个与后一个建立联系
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
	// 头节点的下一个节点尾为节点
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
// 将key提升为最近使用的
// 1、校验是否存在
// 2、如果不存在不需要操作
// 3、删除原有元素
// 4、在尾节点添加
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
	// hash中校验是否存在来保证操作时间复杂度O(1)
	node, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	// 将当前节点提升为最近可用
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
		return
	}

	if this.cap == this.cache.Size() {
		// 删除最近未使用的元素
		this.removeLeastRecently()
	}
	// 添加为最近使用的元素
	this.addRecently(key, value)
}

func (this *LRUCache) Println() {
	if rs, err := json.Marshal(this.hashMap); err == nil {
		log.Println(string(rs))
	}
}
