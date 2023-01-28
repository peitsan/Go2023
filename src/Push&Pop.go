package main

import (
	"errors"
	"fmt"
)

const MAXSIZE = 10

type DataType struct {
	stuName string
	stuNum  string
}

type Stack struct {
	MaxTop int8
	Top    int8
	data   [MAXSIZE]DataType
}

func (this *Stack) Push(val DataType) (err error) {

	//先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据 从栈顶部插入一个元素
	this.data[this.Top] = val
	return
}

func (this *Stack) Pop() (dataType DataType, err error) {
	//先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return dataType, errors.New("stack empty")
	}
	//出栈
	dataType = this.data[this.Top]
	this.Top--
	return dataType, nil
}

func main() {
	stk := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	dataType := DataType{stuName: "Xiaoming", stuNum: "2021212961"}
	stk.Push(dataType)
	fmt.Printf("入栈元素: %v", stk.data[stk.Top])
	pop, _ := stk.Pop()
	fmt.Printf("出栈元素: %v", pop)
}
