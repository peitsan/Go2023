package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=helloc %v 取出表达式的值value
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2} 带key打印
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2} 带类型注释 为最前面定义的结构体 point类型

	f := 3.141592653211292223189262665529
	fmt.Println(f) // 3.141592653
	fmt.Println("%b\n", f)
	fmt.Printf("%.2f\n", f) // 3.14 与python相同
	// 如果值是一个结构体，%+v 变体将包括结构体的字段名。
	fmt.Printf("%+v\n", p) //输出结果为 {x:1 y:2}

	// %#v 变体打印值的 Go 语法表示，即将生成该值的源代码片段。
	fmt.Printf("%#v\n", p) //输出结果为 main.point{x:1, y:2}

	// 打印类型
	fmt.Printf("%T\n", p) //输出结果为 main.point

	// 打印布尔值
	fmt.Printf("%t\n", true) //输出结果为 true

	// 打印整数。
	fmt.Printf("%d\n", 123) //输出结果为 123

	// 打印二进制
	fmt.Printf("%b\n", 14) //输出结果为 1110

	// 打印字符
	fmt.Printf("%c\n", 33) //输出结果为 ！

	// 打印 16 进制
	fmt.Printf("%x\n", 456) //输出结果为 1c8

	// 打印浮点数
	fmt.Printf("%f\n", 78.9) //输出结果为 78.900000
	// 指数型
	fmt.Printf("%e\n", 123400000.0) //输出结果为 1.234000e+08
	fmt.Printf("%E\n", 123400000.0) //输出结果为 1.234000E+08
	// 每个字符用两位 16 进制表示。
	fmt.Printf("%x\n", "hex this") //输出结果为 6865782074686973

	// 打印指针`.
	fmt.Printf("%p\n", &p) //输出结果为 0xc00001c08

	// 字符宽度
	fmt.Printf("|%6d|%6d|\n", 12, 345) //输出结果为 |    12|   345|

	//字符精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //输出结果为 |  1.20|  3.45|

	// 左对齐
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //输出结果为 |1.20  |3.45  |

	//同样可以控制字符的宽度
	fmt.Printf("|%6s|%6s|\n", "foo", "b") //输出结果为 |   foo|     b|

	// 同样字符左对齐.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //输出结果为 |foo   |b     |

	// 合并
	s = fmt.Sprintf("a %s", "string") //输出结果为 a string 这里不能用短声明 上面已经声明了s
	fmt.Println(s)                    //输出结果为
	// 同样的效果
	fmt.Fprintf(os.Stderr, "an %s\n", "error") //输出结果为 an error
}
