package main

import (
	"fmt"
	"goCode/xiecheng/model"
	"goCode/xiecheng/service"
	"time"
)

func main() {

	//协程初学，关键字 go
	go test()
	for i := 0; i <= 10; i++ {

		fmt.Println("main() say hellow", i)
		time.Sleep(time.Second)
	}

	user := model.GetUser("张三", 11)
	name := user.GetName()

	//获取cpu数
	cpuService := service.GetCpuService()
	cpuService.Test()
	fmt.Println(name)
}

func test() {
	for i := 0; i <= 10; i++ {

		fmt.Println("test() say hellow", i)
		time.Sleep(time.Second)
	}
}
