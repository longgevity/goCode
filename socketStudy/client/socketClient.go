package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	con, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}

	fmt.Println("net.Dial succ", con)

	//客户端发送数据给服务端
	reds := bufio.NewReader(os.Stdin) //s.Stdin表示标准输入，即终端
	//从终端读取用户的输入，并发送给服务器
	for {
		mes, err := reds.ReadString('\n')
		if err != nil {
			fmt.Println("reds.ReadString err", err)
			return
		}
		if mes == "exit" {
			return
		}
		n, err := con.Write([]byte(mes))
		if err != nil {
			fmt.Println("con.Write err", err)
			return
		}
		fmt.Println("发送消息 result", n)
	}

}
