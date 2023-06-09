package service

import (
	"fmt"
	"time"
)

type channalDemo struct {
}

func GetChannalDemo() channalDemo {
	return channalDemo{}
}

func writeChan(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		time.Sleep(time.Second)
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

func (channalDemo *channalDemo) Work() {

	intChan := make(chan int, 50)

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
