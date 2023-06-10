package main

import (
	"fmt"
	"goCode/channalDemo/service"
)

func main() {

	//只写管道
	var channaWrite chan<- int
	channaWrite = make(chan<- int, 3)
	channaWrite <- 3
	//只读管道应用

	var chanRead <-chan int
	chanRead = make(<-chan int, 3)
	num3 := <-chanRead
	fmt.Println(num3)
	//定义一个可以存放三个数字的管道

	var channa chan int

	channa = make(chan int, 3)

	//输出channal
	fmt.Printf("channa的值%v chanal的地址是%p\n", channa, &channa)

	//管道写入,当给管道写入数据时，不能超过管道容量
	channa <- 10
	num := 211
	channa <- num

	//输出管道的长度容量
	fmt.Println(len(channa), cap(channa))

	//管道数据取出
	num2 := <-channa
	fmt.Println("num=", num2, len(channa), cap(channa))

	//管道便利关闭
	channalFor := service.GetChannalFor()
	channalFor.ChannalFor()

	//管道、协程结合使用
	// channalDeo := service.GetChannalDemo()
	// channalDeo.Work()

	//求素数
	chanSushu := service.GetChannalSushu()
	chanSushu.Sushu(8000)
}
