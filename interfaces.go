package main

import "fmt"

type email struct {
	isSubscribed bool
	body         string
}
type insta struct {
	isSubscribed bool
	body         string
}

func (e email) cost() float64 {
	if e.isSubscribed {
		return 0.01 * float64(len(e.body))
	} else {
		return 0.05 * float64(len(e.body))
	}
}
func (e insta) cost() float64 {
	if e.isSubscribed {
		return 0.02 * float64(len(e.body))
	} else {
		return 0.1 * float64(len(e.body))
	}
}
func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
	print()
}
