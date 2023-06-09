package service

import (
	"fmt"
)

type channalDemo struct {
}

func GetChannalDemo() channalDemo {
	return channalDemo{}
}

func writeChan(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		// time.Sleep(time.Second)
	}
	//写完关闭管道
	close(intChan)
}

func readChan(intChan chan int, exitChan chan bool) {
	for {
		x, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read", x)
	}
	exitChan <- true
	close(exitChan)
}

/*
*
一个编译器发现一个管道只有写，超出管道容量，会发生阻塞；如果存在读，读的很慢，频率不一致是没有问题的
*/
func (channalDemo *channalDemo) Work() {

	intChan := make(chan int, 1000)

	exitChan := make(chan bool, 1)
	go writeChan(intChan)
	go readChan(intChan, exitChan)

	for {
		x, _ := <-exitChan
		if x {
			break
		}

	}
}

/*启动一个携程写1-2000
启动八个携程取管道里面数据，再放入另外一个管道
*/
