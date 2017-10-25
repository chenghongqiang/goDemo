package main

/*
1.Go程序可以由多个标记组成，可以是关键字，标识量，敞亮，字符串，符号
fmt.Println("hello world") 6个标记 每行一个
1.fmt
2. .
3.Println
4.(
5.hello world
6.)

2.标识符用来命名变量、类型等程序实体，第一个字符必须是字母或下划线而不能是数字

3.Go语言中变量的声明必须用空格隔开

4.不同类型的局部和全局变量默认值为
int 0
float32 0
pointer nil


*/

var x, y int
var ( //这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"  //出现在:=左侧的变量不应该是已经被声明过的，否则会编译错误

func test() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)

}

func main1() {
	const (
		a = iota
		b = iota
		c
	)

	println(a, b, c)

}
