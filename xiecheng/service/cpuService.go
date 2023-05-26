package service

import (
	"fmt"
	"runtime"
)

type CpuService struct {
}

func GetCpuService() CpuService {
	return CpuService{}
}

func (CpuService *CpuService) Test() {
	cpunum := runtime.NumCPU()

	fmt.Println(cpunum)
}

func main() {
	cpuService := GetCpuService()

	cpuService.Test()
}

/*
*
需求：现在要计算 1-200对各个数的阶乘，并且把各个数的阶乘放入到map中
最后现实出来。要求使用goroutine完成
*/
func (CpuService *CpuService) Demo1(map1 map[int]int, n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}
	map1[n] = res
}
