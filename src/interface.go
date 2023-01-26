package main

import "fmt"

type Phone interface {
	speak()
	getPrice()
}

type IPhone struct {
	name  string
	price int32
}

type Oppo struct {
	name  string
	price int
	color string
}

type Mi struct {
	f bool
}

// 手机讲话区域（对相关的接口进行实现）
func (P IPhone) speak() {
	fmt.Println("Hi，我是 Siri！")
}

func (P Oppo) speak() {
	fmt.Println("我是", P.name)
}

func (P Mi) speak() {
	fmt.Println("大家好,我是小爱童鞋!")
}

// 获取价格
func (P IPhone) getPrice() {
	fmt.Println("My price is", P.price)
}

func show(myPhone Phone) {
	myPhone.speak()
	myPhone.getPrice()
}

func main() {
	// 将新建对象传入展示大舞台,大舞台代码不变,展示不同的效果
	myPhone := IPhone{
		name:  "华为13远峰蓝",
		price: 8848,
	}
	show(myPhone)

	var i interface{}
	i = 10
	//没有任何方法的接口就是空接口,实际上每个类型都实现了空接口,所以空接口类型可以接受任何类型的数据
	v, ok := i.(int)
	fmt.Println(v, ok)
	v2 := i.(int)
	fmt.Println(v2)
}
