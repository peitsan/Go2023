package main

import (
	"fmt"
	"sync"
	"time"
)

// 异步协程 sync库
var (
	lock1 = &sync.Mutex{}
	lock2 = &sync.Mutex{}
)

func repeater1() {
	//可以实现一个异步循环的定时器 不会造成内存的重复申请 所有异步请求只在同一块内存使用  通过lock控制 当内存溢出时上锁
	for i := 1; i < 10; i++ { //只打印10s
		time.Sleep(time.Second)
		lock1.Lock()
		lock2.Lock()
		fmt.Print("o")
		fmt.Print("v")
		fmt.Print("e")
		fmt.Print("r")
		fmt.Print("-")
		lock2.Unlock()
		lock1.Unlock()
	}
}

func repeater2() {
	for i := 1; i*i < 100; i++ {
		time.Sleep(time.Minute)
		lock1.Lock()
		lock2.Lock()
		fmt.Print("e")
		fmt.Print("r")
		fmt.Print("o")
		fmt.Print("v")
		fmt.Print("-")
		lock1.Unlock()
		lock2.Unlock()
	}
}
func repeater3() {
	for i := 1; i*i < 100; i++ {
		time.Sleep(2 * time.Minute)
		lock1.Lock()
		lock2.Lock()
		fmt.Print("r")
		fmt.Print("o")
		fmt.Print("e")
		fmt.Print("v")
		fmt.Print("-")
		lock1.Unlock()
		lock2.Unlock()
	}
}
func repeater4() {
	for i := 1; i*i < 100; i++ {
		time.Sleep(2 * time.Minute)
		lock1.Lock()
		lock2.Lock()
		fmt.Print("e")
		fmt.Print("r")
		fmt.Print("o")
		fmt.Print("v")
		fmt.Print("-")
		lock1.Unlock()
		lock2.Unlock()
	}
}
func main() {
	//repeater1()
	//repeater2()
	//	遵循严格异步执行  repeater1执行完成后才执行 repeater2
	//	over-over-over-over-over-over-over-over-over-erov-erov-erov-erov-erov-erov-erov-erov-erov-

	go repeater1() //线程长度  为1  先进入goroutie 形成线程队列  线程比线程总长度短 先执行 每s打印over-over-
	go repeater3() //线程长度  为2s 线程=线程总长度 同步执行
	go repeater4() //线程长度  为2s 线程=线程总长度 同步执行
	go repeater1() //线程长度  为1s 先进入goroutie 形成线程队列  线程比线程总长度短 先执行 每s打印over-over-
	go repeater2()
	go repeater3() //线程长度  为2s 线程=线程总长度 同步执行
	go repeater4() //线程长度  为2s 线程=线程总长度 同步执行
	repeater3()    //线程长度  为2s 线程=线程总长度 同步执行
	//go : goroutine
	//goroutine是一个并发编程关键字 函数调用时是并发执行的，goroutine默认使用电脑的所有核心，其他核在进行go关键字后面的函数运算，但是由于程序执行完main函数结束后立刻结束，因此不确定其他核的函数有没有执行完毕，也就造成了每次运行打印的数字个数不一样。
	//	解析： 4句goroutine使 原本的异步调用定义形成并发对列 最后与 repeater1的异步打印同步执行
	//	解析： 当异步线程长度不等时  相同线程长度的程序会同步执行 线程短的先执行 线程长的后执行 当线程耗散时若长线程进程仍未执行会跳出进程线程
	// 线程长度取决于 goroutine 外的第一个异步线程
}
