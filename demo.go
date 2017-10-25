package main

import (
	"fmt";
	"os";
	"io"
)

func main() {

	a  := []int {1,2,5}
	fmt.Println(a)

	reverse(a)
	fmt.Println(a)

	test()
}

func reverse(a []int) {

	for i,j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

}

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer f.Close() //会在函数结束后运行

	var result []byte
	buf := make([]byte, 100)

	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}	
			return "", err
		}
	}

	return string(result), nil


}

//defer 语句用于预设一个函数调用（即推迟执行函数），该函数会在执行deferd的函数返回之前立即执行）
//可用来处理无论以何种路径返回，都必须释放资源的函数，典型例子是解锁互斥和关闭文件
//这样的函数调用有两点好处 1.能确保你不会忘记关闭文件 2.意味着关闭离打开很近，这总比放在函数结尾处清晰明了

//被推迟函数的实参在推迟执行时就会求值，而不是在调用执行时求值 无需担心变量在函数执行过程中被改变，同时意味着
//单个已推迟的调用可推迟多个函数的执行

func test() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	
}
//打印 4 3 2 1 0，被推迟的函数按照后进先出LIFO


