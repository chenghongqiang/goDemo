package main

func main_logic() {
	var a int = 1
	var b int = 2
	var ret = max(a, b)

	println("max value is:", ret)

}

func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
