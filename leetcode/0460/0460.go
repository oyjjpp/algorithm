package leetcode

import (
	"container/list"
)

type Node struct {
	key  int
	val  int
	freq int
}

type LFUCache struct {
	// key 到val的映射
	keyToVal map[int]*list.Element
	// freq到key列表的映射
	freqTokeys map[int]*list.List
	minFreq    int
	cap        int
}

// Constructor
// 初始化函数
func Constructor(capacity int) LFUCache {
	return LFUCache{
		keyToVal:   map[int]*list.Element{},
		freqTokeys: map[int]*list.List{},
		minFreq:    0,
		cap:        capacity,
	}
}

// Get
// 获取值
func (this *LFUCache) Get(key int) int {
	// 不存在时则返回-1
	ele, ok := this.keyToVal[key]
	if !ok {
		return -1
	}
	// 增加key对应的freq
	return this.updateFreq(ele)
}

// updateFreq
// 访问key的同时需要增加对应freq的次数
func (this *LFUCache) updateFreq(ele *list.Element) int {
	data := ele.Value.(*Node)
	// 处理freq to key的映射关系
	curList, ok := this.freqTokeys[data.freq]
	if !ok {
		return -1
	}
	curList.Remove(ele)

	// TODO 确认是否需要更新
	if curList.Len() == 0 {
		if data.freq == this.minFreq {
			this.minFreq++
		}
	}
	// 更新key与freq的关系，当key对应的freq增加时候，则存储到对应的链表中
	data.freq++
	newList, ok := this.freqTokeys[data.freq]
	// 如果不存在链表则创建
	if !ok {
		newList = list.New()
		this.freqTokeys[data.freq] = newList
	}
	newEle := newList.PushBack(data)

	// 更新元素的freq
	this.keyToVal[data.key] = newEle
	return data.val
}

// Put
// 存入新的值
func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	// 如果key已存在，修改对应的val即可
	if ele, ok := this.keyToVal[key]; ok {
		data := ele.Value.(*Node)
		data.val = value
		// 调整key的freq次数
		this.updateFreq(ele)
		return
	}

	// key不存在则需要插入
	// 容量已满的话需要淘汰一个frqp最小的key
	if this.cap == len(this.keyToVal) {
		this.removeMinFreqKey()
	}

	// 创建新的节点
	newNode := &Node{
		key:  key,
		val:  value,
		freq: 1,
	}
	// 更新freq to key的链表
	newList, ok := this.freqTokeys[newNode.freq]
	if !ok {
		// 不存在则创建一个链表
		newList = list.New()
		this.freqTokeys[newNode.freq] = newList
	}
	newEle := newList.PushBack(newNode)
	this.keyToVal[key] = newEle
	this.minFreq = newNode.freq
}

// removeMinFreqKey
// 淘汰一个freq最小的key
func (this *LFUCache) removeMinFreqKey() {
	minList, ok := this.freqTokeys[this.minFreq]
	if !ok {
		return
	}
	delE := minList.Front()
	minList.Remove(delE)
	data := delE.Value.(*Node)
	delete(this.keyToVal, data.key)
}
