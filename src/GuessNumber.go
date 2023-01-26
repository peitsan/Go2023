package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())  //获取一个随机种子
	secretNumber := rand.Intn(maxNum) //rand一个不超过3位数的整数
	// 作弊模式
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	// 通过一个 for 循环实现一直猜数，直到猜中
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess) //键盘输入 fmt.Scanf 赋值给 guess 变量
		if err != nil {
			// 匹配错误报错了 打印提属于
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
