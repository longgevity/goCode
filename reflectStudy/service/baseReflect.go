package service

import (
	"fmt"
	"reflect"
)

type baseReflectService struct {
}

func GetBaseReflectService() baseReflectService {
	return baseReflectService{}
}

/*
设置interface，可接受所有参数
*/
func reflectTest(i interface{}) {
	//获取type数据信息（reflect.type)
	iType := reflect.TypeOf(i)
	fmt.Println("i的type：=", iType)
	//获取value（reflect.value）
	iValue := reflect.ValueOf(i)

	//修改值
	iValue.Elem().SetInt(200)
	//iValue转interface
	// endi := iValue.Interface()
	//将interface通过断言转换成需要的类型
	// num2 := endi.(int)
	// fmt.Println(num2)
}

func (baseReflectService *baseReflectService) BaseRun() {

	var i int = 100
	reflectTest(&i)

	fmt.Println("主程序i", i)
}
