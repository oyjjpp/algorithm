package leetcode

type LFUCache struct {
	// key 到val的映射
	keyToVal map[int]int
	// key到freq的映射
	keyToFreq map[int]int
	// freq到key列表的映射
	freqTokeys map[int][]int
	minFreq    int
	cap        int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keyToVal:   map[int]int{},
		keyToFreq:  map[int]int{},
		freqTokeys: map[int][]int{},
		minFreq:    0,
		cap:        capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	// 不存在时则返回-1
	data, ok := this.keyToFreq[key]
	if !ok {
		return -1
	}
	// 增加key对应的freq
	return data
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	// 如果key已存在，修改对应的val即可
	if _, ok := this.keyToVal[key]; ok {
		// 更改key的值
		this.keyToVal[key] = value
		// 调整key的freq次数
		this.increaseFreq(key)
		return
	}

	// key不存在则需要插入
	// 容量已满的话需要淘汰一个frqp最小的key
	if this.cap <= len(this.keyToVal) {
		this.removeMinFreqKey()
	}

	// 插入key和val，对应的freq为1
	this.keyToVal[key] = value
	this.keyToFreq[key] = 1

	// 更新freq到key列表的映射
	if data, ok := this.freqTokeys[1]; ok {
		data = append(data, key)
	} else {
		this.freqTokeys[1] = []int{key}
	}

	// 插入新的key后最小分freq肯定是1
	this.minFreq = 1
}

// increaseFreq
// 增加对应key的freq
func (this *LFUCache) increaseFreq(key int) {
	// 更新key to freq
	freq := this.keyToFreq[key]
	this.keyToFreq[key] = freq + 1

	// 更新freq to key
}

// removeMinFreqKey
// 淘汰一个freq最小的key
func (this *LFUCache) removeMinFreqKey() {
	// freq最小的key列表
	keyList := this.freqTokeys[this.minFreq]

	// 最先被插入的key就是应该被淘汰的key
	deleteKey := keyList[0]
	if len(keyList) == 1 {
		delete(this.freqTokeys, this.minFreq)
	} else {
		// TODO 待确定是否能修改原来切片
		keyList = keyList[1:]
	}
	delete(this.keyToFreq, deleteKey)
	delete(this.keyToVal, deleteKey)
}
