package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
)

// 这部分是线程不安全的存在

//var buffer bytes.Buffer
//var wg sync.WaitGroup

//func main() {
//	runtime.GOMAXPROCS(runtime.NumCPU())

//	wg.Add(2)

//	go func() {
//		for i := 0; i < 1000; i++ {
//			buffer.WriteString("a")
//		}

//		wg.Done()
//	}()

//	go func() {
//		for i := 0; i < 1000; i++ {
//			buffer.WriteString("b")
//		}

//		wg.Done()
//	}()

//	wg.Wait()

//	fmt.Println(buffer.String())
//}

// 将 bytes.Buffer 与 锁包裹在一起，就是简单的线程安全的了

type Buffer struct {
	b  bytes.Buffer
	rw sync.RWMutex
}

// 非写状态时, 多个线程可以同时进行读操作.
// 写状态是, 其他线程既不能读也不能写, 性能比单纯的 Mutex 好
func (b *Buffer) Read(p []byte) (n int, err error) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	return b.b.Read(p)
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.rw.Lock()
	defer b.rw.Unlock()

	return b.b.Write(p)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	buffer := Buffer{}
	a := []byte{'a'}
	b := []byte{'b'}
	r := make([]byte, 2000)

	go func() {
		for i := 0; i < 1000; i++ {
			buffer.Write(a)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			buffer.Write(b)
		}

		wg.Done()
	}()

	wg.Wait()

	buffer.Read(r)
	fmt.Println(string(r))
}
