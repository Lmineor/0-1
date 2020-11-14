package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp demo客户端
func main() {
	// 1. 与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Dail failed, err is %v", err)
		return
	}

	// 2. 利用该连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		// 给服务端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("Error is %v\n", err)
			return
		}
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed, err: %v\n", err)
			return
		}
		fmt.Printf("收到服务端回复: %v\n", string(buf[:n]))
	}
}
