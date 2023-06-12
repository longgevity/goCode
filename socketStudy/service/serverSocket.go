package main

import (
	"fmt"
	"net"
)

func receiveMes(con net.Conn) {
	defer con.Close()
	//循环接收客户端消息
	for {
		//创建一个切片接受消息
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息
		//如果客户端不发送消息，则协程阻塞
		//fmt.Printf("wait client %v message\n", con.RemoteAddr())
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("con.Read err", err)
			return
		}

		//显示客户端发送的消息内容到服务器终端
		fmt.Print("receive message:=", string(buf[:n]))

	}
}

func main() {

	fmt.Println("服务器端启动，开启监听。。。")
	//tcp表示使用的协议，0.0.0.0:8888表示监听本地8888端口
	lis, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听端口失败err", err)
		return
	}
	defer lis.Close()
	fmt.Println("listen = ", lis)

	//循环等待客户端等待
	for {
		//等待联结
		fmt.Println("等待链接")
		con, err := lis.Accept()
		if err != nil {
			fmt.Println("链接失败")
		} else {
			fmt.Println("con =", con)
			fmt.Printf("远程连接的ip%v", con.RemoteAddr())
		}
		go receiveMes(con)
	}

}
