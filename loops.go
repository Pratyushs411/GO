package main

import "fmt"

//for initial;condition;after{}

func printTo10() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Print("\n")
}

func bulkMessages(n int) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += 1.0 + 0.01*(float64(i))
	}
	return sum
}

// while loop
func sumUptoN(n int) {
	i := 1
	sum := 0
	for i <= n {
		sum += i
		i++
	}
	fmt.Printf("\nsum upto %v is :%v\n", n, sum)
}
