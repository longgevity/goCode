package service

import "fmt"

type chanSushu struct {
}

func GetChannalSushu() chanSushu {

	return chanSushu{}
}

func addChan(chanInt chan int, num int) {

	for i := 1; i <= num; i++ {
		chanInt <- i
	}
	close(chanInt)
}

func addSushu(chanInt chan int, resultChan chan int, exitChan chan bool) {

	for {
		num, ok := <-chanInt
		if ok {
			if is_prime(num) {
				resultChan <- num
			}
		} else {
			break
		}
	}
	exitChan <- true
}

func is_prime(num int) bool {
	//从2遍历到n-1，看看是否有因子
	for i := 2; i < num; i++ {
		if num%i == 0 {
			//发现一个因子被整除
			return false
		}
	}
	return true
}

/*
多协程获取数字以内的素数，并且展示出来
*/
func (chanSushu *chanSushu) Sushu(num int) {

	chanInt := make(chan int, 100)

	resultChan := make(chan int, 500)

	exitChan := make(chan bool, 4)

	go addChan(chanInt, num)

	for i := 0; i < 4; i++ {
		go addSushu(chanInt, resultChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
			fmt.Println(i, "协程执行完成")
		}
		fmt.Println("素数获取完毕")
		close(resultChan)
	}()

	for v := range resultChan {
		fmt.Println("获取到素数{}", v)
	}

	fmt.Println("退出主线程")

}
