package main

import "fmt"

func main() {
	var map1 map[string]string

	map1 = make(map[string]string)

	map1["key"] = "value"

	fmt.Println(map1)

}
