package service

type reflectTwoExerciseService struct {
}

func GetReflectTwoExerciseService() reflectExerciseService {
	return reflectExerciseService{}
}

type StudentTwoStruce struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int
}
