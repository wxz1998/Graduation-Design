package main

import (
	"fmt"
	"reflect"
)

/*
 * @Author: zut.wxz
 * @Date: 2020-12-02 15:12:58
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-02 16:21:57
 * @Description:
 */

// 在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}

type name和type kind
在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，
但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。
举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct
}