package main

import "fmt"

func main() {
	//并不是遍历 是一个键值对字典
	m := make(map[string]int) //构造一个索引为string类型变量的 键值对为 int类的字典
	m["one"] = 1              //插入一个键值对
	m["two"] = 2              //插入第二个键值对
	fmt.Println(m)            // map[one:1 two:2]
	fmt.Println(len(m))       // 2
	fmt.Println(m["one"])     // 1
	fmt.Println(m["unknown"]) // 0 找不到

	r, ok := m["unknown"] //短声明  m["unknown"] = 0         等价于r, ok := 0
	fmt.Println(r, ok)    // 0 false ?   A变量,ok 表示A变量是否存在 存在返回true  反之返回false

	delete(m, "one") //删除 键为one 的键值对
	fmt.Println(m)   // map[ two:2]

	m2 := map[string]int{"one": 1, "two": 2} //支持短声明和普通声明 支持解包声明
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)

	//Comma-ok断言的语法是：value, ok := element.(T)。element必须是接口类型的变量，T是普通类型。如果断言失败，ok为false，否则ok为true并且value为变量的值。来看个例子：
	a := "string"
	fmt.Println(a, ok)
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Print(value, ",", ok)
}
