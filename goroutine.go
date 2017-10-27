package main

import (
	"fmt"
	"time"
)

func loop() {
	for i :=0 ;i<10; i++ {
		fmt.Printf("%d ", i)		
	}
}

func channel() {
	var messages chan string = make(chan string)
	func(message string) {
		messages <- message  //存消息
	}("ping")

	fmt.Println(<- messages) //取消息

}

func main() {
	_  = time.Date
	channel()
}

