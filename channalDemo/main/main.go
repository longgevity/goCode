package main

import "fmt"

func main() {
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
}
