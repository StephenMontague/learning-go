package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println(mySlice)

	sort.Ints(mySlice)

	log.Println(mySlice)
}