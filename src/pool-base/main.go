package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 给予 New 函数, 在数据不足的情况下会自动创建
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	// 不给予自动创建的 New 函数, 需要手动添加元素到链表中
	//	p := &sync.Pool{}
	//	p.Put(0)

	a := p.Get().(int)
	p.Put(1)

	// pool 在 init 的时候就注册了一个 poolCleanup 函数, 在每次 GC 之前都会调用, 清楚所有 pool 内缓存的对象
	// pool 内对象的生存周期只限于两次 GC 之间的这段时间
	// 正因为这样, Pool 并不适用于类似 socket 连接池的使用
	runtime.GC()

	b := p.Get().(int)
	fmt.Println(a, b)
}
