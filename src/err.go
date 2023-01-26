package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

// 声明一个校验函数findUser 传入参数为 一个 user类型 的对象数组 和一个string 类型的名字   返回值定义 若有返回值 返回 user类型的指针 v 若无返回 输出error类型 的报错
func findUser(users []user, name string) (v *user, err error) {
	//遍历输入的用户数据数组 user
	for _, u := range users {
		//如果查询到 {name}的 人 则由返回值 通过 地址带回主函数 没哟报错信息
		if u.name == name {
			return &u, nil
		}
	}
	//如果没有查询到 {name}的 人 则返回值v* user 为 无值 nil 带回主函数 报错信息 new一个 notfound
	return nil, errors.New("not found")
}

// nil(non-value)是无值类型 与C++里的NULL不同，NULL是一个宏定义，NULL值为0，nil表示无值
func main() {
	u, err := findUser([]user{{"wang", "1024"}, {"yuan", "1212"}}, "wang")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name, u.password) // wang

	databaseSearchRes, err2 := findUser([]user{{"iiru", "3096"}, {"ok", "5858"}, {"wang", "1024"}, {"yuan", "1212"}}, "iiru")
	if err2 != nil {
		//如果有错误 打印错误信息
		fmt.Println(err2)
	}
	fmt.Println(databaseSearchRes.name, databaseSearchRes.password)
	//和上面分开的写法是一样的 逗号表达式  xxx, 条件 {}  xxx先执行赋值语句
	if u, err := findUser([]user{{"wang", "1024"}, {"yuan", "1212"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}
