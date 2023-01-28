package main

import (
	"fmt"
	"time"
)

func main() {
	go learnFrontend()
	go learnAndroid()
	go learnMachineLearning()
	learnMachineLearning()
}

func learnBackend() {
	time.Sleep(10 * time.Second) // 十分钟速通 web 后端（确信
	fmt.Println("会了！10Second")
}

func learnFrontend() {
	time.Sleep(time.Nanosecond)
	fmt.Println("会了！Nanosecond")
}

func learnAndroid() {
	time.Sleep(20 * time.Second)
	fmt.Println("悔了！20Second")
}

func learnMachineLearning() {
	time.Sleep(114 * time.Second)
	fmt.Println("废了！114Second")
}
