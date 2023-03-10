package main

import (
	"fmt"
	"time"
)

type Item struct {
	Name  string
	Count int
}

func main() {
	//通道(chan)类似于一个队列，特性就是先进先出，多用于goruntine之间的通信
	pipeline := make(chan Item, 10) // 一条可以放 10 个 item 的流水线
	go func() {
		for {
			time.Sleep(1 * time.Second)
			pipeline <- Item{
				"螺丝",
				5,
			} //生产螺丝耗时1s 耗费 5c
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			pipeline <- Item{
				Name:  "齿轮",
				Count: 3,
			} //生产螺丝耗时2s 耗费 3c
		}
	}()
	for {
		item := <-pipeline
		fmt.Printf("%#v\n", item)
	}
}
