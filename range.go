package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		fmt.Println("for num:", num)
		sum += num
	}
	fmt.Println("sum:", sum)

	//在数组上使用range将传入index和值两个变量。上面例子我们不需要使用该元素的符号，所以使用空白符_
	//省略了，有时候我们需要索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range也可以用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串，第一个参数是字符的索引，第二个是字符Unicode值本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
