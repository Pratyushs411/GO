package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}
func temp(temp1 func(int, int) int, y int) int {
	return temp1(2, 3) + y
}
func names() (string, string) {
	return "pratyush", "aarush"
}

func main() {
	//email := "pratyid.com"
	// if leng := len(email); leng < 20 {
	// 	fmt.Println("good")
	// } else {
	// 	fmt.Println("not good")
	// }
	fmt.Println(add(2, 3))
	fmt.Println(temp(add, 4))

	x, _ := names()
	fmt.Println(x)
	fmt.Println(pratyush.age)
	fmt.Println(r.area())
	fmt.Println(t.area())
	e := email{
		isSubscribed: true,
		body:         "this is subscribed",
	}
	ins := insta{
		isSubscribed: true,
		body:         "this is subscribed",
	}
	fmt.Println(e.cost())
	fmt.Println(ins.cost())
}
