package main

import "goCode/reflectStudy/service"

func main() {

	//基本数据类型使用
	// baseService := service.GetBaseReflectService()
	// baseService.BaseRun()
	//对象使用
	// structService := service.GetStructReflectService()
	// structService.StruceRun()

	//对象字段、方法调用
	reflectExerciseService := service.GetReflectExerciseService()
	reflectExerciseService.RunService()
}

/*
注意事项：
1、reflect.Value.kind,获取变量的类别，返回的是一个常量
2、type是类型,kind是类别，Type跟kind可能是相同的，也可能是不同的；如果int、sclice等相同，struct对象则不同
3、通过反射可以让interface{}和Refelct.value之间相互转换
4、使用反射获取变量的值（并且返回对应的类型），要求数据类型匹配，比如x是int，那么就应该使用reflect,Value(x).Int(),而不能使用其他的
5、修改值 iValue.Elem().set***(),传参必须是指针类型数据
*/
