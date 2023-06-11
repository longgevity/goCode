package service

import (
	"fmt"
	"reflect"
)

type reflectExerciseService struct {
}

func GetReflectExerciseService() reflectExerciseService {
	return reflectExerciseService{}
}

type StudentStruce struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int
}

func (stu *StudentStruce) GetName() string {

	return stu.Name
}

func (stu *StudentStruce) SetName(name string) {

	stu.Name = name
}

func (stu *StudentStruce) MakeStu(name string, age, sex int) {
	stu.Name = name
	stu.Age = age
	stu.Sex = sex
}

func TestStruct(a interface{}) {
	rType := reflect.TypeOf(a)
	rValue := reflect.ValueOf(a)

	kd := rValue.Elem().Kind()
	fmt.Println("rType=", rType)
	fmt.Println("rValue=", rValue)
	if kd != reflect.Struct {
		fmt.Println("非实体数据类型")
		return
	}

	//获取结构体下所有字段
	fileNum := rValue.Elem().NumField()
	fmt.Printf("结构体有%v字段\n", fileNum)
	for i := 0; i < fileNum; i++ {
		fmt.Printf("file%v值为 ： %v \n", i, rValue.Elem().Field(i))
		//获取到struct标签，主意事用reflect.Type来获取tag标签的值
		tagVal := rType.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("tag%vtag为 ： %v \n", i, tagVal)
		}
	}

	//获取结构体下有多少方法
	methodNum := rValue.NumMethod()
	fmt.Printf("结构体有%v方法\n", methodNum)
	for i := 0; i < methodNum; i++ {
		fmt.Printf("method%v的名称%v\n", i, rType.Method(i).Name)

	}
	//结构题方法的回掉使用call(),方法的排序默认是按照 函数名的排序（ASCII码）
	// rValue.Method(0).Call(nil) //获取第一个方法reflect.value.Method(i)

	var param []reflect.Value //声明穿参[]reflect.Value
	param = append(param, reflect.ValueOf("李四"))
	param = append(param, reflect.ValueOf(100))
	param = append(param, reflect.ValueOf(0))

	res := rValue.Method(1).Call(param) //传入参数[]reflect.Value,返回[]reflect.Value
	fmt.Println("res:=", res)

	name := rValue.Method(0).Call(nil)
	fmt.Println("获取名字", name[0].String())

	var param1 []reflect.Value //声明穿参[]reflect.Value
	param1 = append(param1, reflect.ValueOf("李四"))
	res1 := rValue.Method(2).Call(param1)
	fmt.Println("获取名字", res1)

}

/*
使用反射遍历结构体字段，调用结构体的方法，并且获取结构题的标签值
*/
func (reflectExerciseService *reflectExerciseService) RunService() {
	stu := StudentStruce{"张三", 1, 2}
	// stu.MakeStu("张三", 20, 1)
	// fmt.Println("初始化的学生", stu)
	fmt.Println("初始化学生信息", stu)
	TestStruct(&stu)
	fmt.Println("执行后学生信息", stu)
}
