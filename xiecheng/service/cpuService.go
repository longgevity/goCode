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
