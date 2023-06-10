package service

import (
	"fmt"
	"reflect"
)

type structReflectService struct {
}

func GetStructReflectService() structReflectService {
	return structReflectService{}
}

type student struct {
	name string
	age  int
}

type animal struct {
	name string
	age  int
}

/*
设置interface，可接受所有参数
*/
func reflectStructTest(i interface{}) {
	//获取type数据信息（reflect.type)
	iType := reflect.TypeOf(i)
	fmt.Println("i的type：=", iType)
	//获取value（reflect.value）
	iValue := reflect.ValueOf(i)
	fmt.Println("i的value ：=", iValue)

	//iValue转interface
	endi := iValue.Interface()
	//将interface通过断言转换成需要的类型
	stu := endi.(student)
	stu.name = "李四"
	fmt.Println("学生姓名：", stu.name, "学生年纪", stu.age)
}

func (structReflectService *structReflectService) StruceRun() {

	stu := student{
		name: "张三",
		age:  11,
	}

	reflectStructTest(stu)

	fmt.Println("主程序i", stu.name)
}
