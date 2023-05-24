package model

type User struct {
	Name string
	Age  int
}

func GetUser(name string, age int) User {

	return User{Name: name,
		Age: age}
}

func (user *User) GetName() string {
	return user.Name
}
