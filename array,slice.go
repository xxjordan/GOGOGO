package main

import "fmt"

// 数组
/*  数组是由固定长度的特定类型元素组成的序列，也就是说，数组的长度也是数组的组成部分。

因为数组的长度是数组类型的一个部分，不同长度或不同类型的数据组成的数组都是不同的类型，因此在Go语言
中很少直接使用数组（不同长度的数组因为类型不同无法直接赋值）。

和数组对应的类型是切片，切片是可以动态增长和收缩的序列，切片的功能也更加灵活，和python中的list类似。
*/

// 以下为数组的几种定义方式
func init_array_1() {
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组,和map[int]Type类型的初始化语法类似, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

/*
go语音中数组变量即表示整个数组，它并不是隐式的指向第一个元素的指针，而是一个完整的值，当一个数组的变量被赋值或者
被传递时，它实际上会复制整个数组。如果数组较大的话，数组赋值也会有较大开销。为了避免，可以传递一个指向数组的指针，
但是数组指针并不是数组。
*/
func init_array_2() {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针

	fmt.Println(a[0], a[1]) // 打印数组的前2个元素
	fmt.Println(b[0], b[1]) // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
}

func main() {
	fmt.Println("init_array_1")
	init_array_1()
	fmt.Println("init_array_2")
	init_array_2()
}
