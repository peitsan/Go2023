package main

import "fmt"

func main() {
	//用make函数 构造一个队列
	que1 := make(chan string, 10)
	//入队
	que1 <- "Xiao"
	que1 <- "Liu"
	que1 <- "Zhao"

	var stack []string
	for {
		q := <-que1              //出队一个元素暂赋给q
		stack = append(stack, q) //入栈一个元素
		if len(que1) == 0 {
			break
		}
	}
	fmt.Println("队列清空，打印输出栈")
	for key, value := range stack {
		fmt.Printf("{%v:%v}", key, value)
	}
}
