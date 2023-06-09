package service

import "fmt"

type channalFor struct {
}

func GetChannalFor() channalFor {

	return channalFor{}
}

func (channalFor *channalFor) ChannalFor() {
	intchan := make(chan int, 3)
	intchan <- 3
	intchan <- 4
	close(intchan)
	//不关闭管道会导致死锁
	for intValue := range intchan {
		fmt.Println("intValue+", intValue)
		fmt.Println("管道长度", len(intchan))
	}

}
