package main

import (
	"fmt"
	"goCode/xiecheng/service"
	"time"
)

func main() {
	cpuService := service.GetCpuService()
	//协程初学，关键字 go
	/*go test()
	for i := 0; i <= 10; i++ {

		fmt.Println("main() say hellow", i)
		time.Sleep(time.Second)
	}

	user := model.GetUser("张三", 11)
	name := user.GetName()*/

	//获取cpu数
	/*cpuService := service.GetCpuService()
	cpuService.Test()
	fmt.Println(name)*/

	//多协程调用
	mapq := make(map[int]int)
	for i := 1; i <= 200; i++ {
		go cpuService.Demo1(mapq, i)
	}

	fmt.Println(mapq)
}

func test() {
	for i := 0; i <= 10; i++ {

		fmt.Println("test() say hellow", i)
		time.Sleep(time.Second)
	}
}
