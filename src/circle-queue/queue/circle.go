package queue

import "sync"

/**
 * 最大数组长度
 * 如果超过该长度，则会返回失败
 */
const MaxLenght uint32 = 100000

/**
 * 基础数据结构
 * head 是一个指向环形队列表头的一个数组下标
 * tail 是一个指向环形队列表尾的一个数组下标
 * length 是一个当前队列内有数据总数
 * value 是存储队列内容的
 */
type CircleQueue struct {
	head uint32
	tail uint32
	length uint32
	lock sync.Mutex
	value[MaxLenght] interface{}
}

/**
 * 检查循环队列是否已满
 */
func (d *CircleQueue) IsFull() bool {
	if (d.length == MaxLenght) {
		return true
	}

	return false
}

/**
 * 查看队列是否是空的
 */
func (d *CircleQueue) IsEmpty() bool {
	if (d.length == 0) {
		return true
	}

	return false
}

/**
 * 存储数据
 * 当数据存满了，则会无法继续存入
 * 返回两个值
 * 第一个值是布尔型，告诉你存储是否成功
 * 第二个值是无符号整型，告诉你当前队列总共有多少数据
 */
func (d *CircleQueue) Put(val interface{}) (bool, uint32) {
	defer d.lock.Unlock()

	d.lock.Lock()

	if (d.IsFull()) {
		return false, d.length
	}

	d.value[d.tail] = val
	d.tail++
	d.length++

	if d.tail % MaxLenght == 0 {
		d.tail = 0
	}

	return true, d.length
}

/**
 * 从队列中获取一个值
 * 返回两个值
 * 第一个值是布尔型，告诉你是否获取成功
 * 第二个值是万能接口，返回具体的值
 */
func (d *CircleQueue) Get() (bool, interface{}) {
	defer d.lock.Unlock()

	d.lock.Lock()

	if (d.IsEmpty()) {
		return false, nil
	}

	val := d.value[d.head]
	d.head++
	d.length--

	if d.head % MaxLenght == 0 {
		d.head = 0
	}

	return true, val
}