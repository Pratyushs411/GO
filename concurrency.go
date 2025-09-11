package main

import (
	"fmt"
	"time"
)

func messagePrint(message string) {
	go func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("message get :%v\n", message)
	}()
	fmt.Printf("message sent:%v\n", message)
	time.Sleep(time.Second)
}
