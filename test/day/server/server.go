package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
监听端口
接收客户端请求建立链接
创建goroutine处理链接。
*/
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recstr := string(buf[:n])
		fmt.Println("收到client端发来的消息:", recstr)
		conn.Write([]byte(recstr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listen from client failed,err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept failed,err:", err)
			continue
		}
		go process(conn)
	}
}
