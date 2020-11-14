package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 处理完要关闭连接
	// 针对当前的连接做数据发送接收操作
	for {
		reader := bufio.NewReader(conn)
		var buff [128]byte
		n, err := reader.Read(buff[:]) // buff[:]做一个切片
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recv := string(buff[:n])
		fmt.Printf("接收到的数据:%v\n", recv)
		conn.Write([]byte("ok"))

	}
}

// tcp server demo
func main() {
	// 1. 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Listen failed, error is %v\n", err)
		return
	}
	for {
		// 2. 等待客户机建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept failed, error is %v\n", err)
			continue
		}
		/// 3. 启动一个goroutine对该连接进行处理
		go process(conn)
	}
}
