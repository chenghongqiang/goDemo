package main

import "fmt"

func Factorial(n int64) (result int64) {
	if n>0 {
		result = n*Factorial(n-1)
		return result
	}

	return 1
} 

func Fibonacci(n int) int{
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(int64(i)))

	var j int
	for j = 0; j < 5; j++ {
		fmt.Printf("斐波拉切数列:%d\n", Fibonacci(j))
	}
	
}