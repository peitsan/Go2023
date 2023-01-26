package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()                                   //获取当前时间的 计算机当前时间 自动转化为东八区
	fmt.Println(now)                                    // 2022-03-27 18:04:59.433297 +0800 CST m=+0.000087933
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC) //将数值型转换为UTC时间戳型
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	fmt.Println(t)                                                  // 2022-03-27 01:25:36 +0000 UTC
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 March 27 1 25 取时间戳
	fmt.Println(t.Format("2006-01-02 15"))                          // 2022-03-27 01:25:36 将t时间戳格式化输出 格式按照 "2006-01-02 15:04:05"这样的 格式 格式可以任意调整
	diff := t2.Sub(t)                                               //t2 - t 的时间差
	fmt.Println(diff)                                               // 1h5m0s
	fmt.Println(diff.Minutes(), diff.Seconds())                     // 65 3900  //单位换算
	t3, err := time.Parse("2006-01-02 15:02:06", "2022-03-27 01:25:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)    // true
	fmt.Println(now.Unix()) // 1648738080
}
