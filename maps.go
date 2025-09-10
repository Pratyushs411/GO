package main

import "fmt"

var ages = make(map[string]int)

func getMap(names []string, ages []int) map[string]int {
	myMap := make(map[string]int)
	i := 0
	for i < len(names) {
		myMap[names[i]] = ages[i]
		i++
	}
	return myMap
}
func printMap(myMap map[string]int) {
	fmt.Println(myMap)
}
