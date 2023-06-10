package main

import "goCode/reflectStudy/service"

func main() {

	//基本数据类型使用
	baseService := service.GetBaseReflectService()
	baseService.BaseRun()
	//对象使用
	structService := service.GetStructReflectService()
	structService.StruceRun()
}
