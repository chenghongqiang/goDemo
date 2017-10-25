package main

import "fmt"

func main() {
	numbers1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers1)

	fmt.Println("numbers == ", numbers1)
	/* 打印子切片从索引1到索引4 4不包含 */
	fmt.Println("numbers[1:4] == ", numbers1[1:4])
	/* 默认下限为0*/
	fmt.Println("numbers1[:3] == ", numbers1[:3])
	/* 默认上限为len(s) */
	fmt.Println("numbers1[2:] == ", numbers1[2:])

	numbers2 := numbers1[1:5]
	printSlice(numbers2)

	numbers1 = append(numbers1, 9)
	printSlice(numbers1)

	numbers1 = append(numbers1, 10)
	printSlice(numbers1)

	numbers3 := make([]int, len(numbers1), (cap(numbers1) * 2))

	copy(numbers3, numbers1)

	printSlice(numbers3)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
