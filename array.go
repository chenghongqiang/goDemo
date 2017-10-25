package main

import "fmt"

func main() {
	var balance = [5]float32{1, 2, 3, 4, 6.0}

	println(balance[4])
	testArray()
}

func testArray() {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
