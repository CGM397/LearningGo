package main

import (
	"fmt"
)

func main() {
	//at least 32 bit in size
	var a int = 10
	fmt.Print(a, " Type: ")
	fmt.Printf("%T\n", a)

	//交由编译器(?)判断类型
	var b = "10"
	fmt.Print(b, " Type: ")
	fmt.Printf("%T\n", b)

	//交由编译器(?)判断类型，省去var关键字
	//声明加赋值
	c := false
	fmt.Print(c, " Type: ")
	fmt.Printf("%T\n", c)

	//0
	var d int
	//""
	var e string
	//false
	var f bool
	//nil
	var g *int
	//[]
	var h []int
	//map[]
	var i map[string]int
	//nil
	var j chan int
	fmt.Println(d, e, f, g, h, i, j)

	//内存地址
	fmt.Println(&a)

	//常量
	const k string = "abc"
	fmt.Println(k)
}
