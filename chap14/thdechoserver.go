package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := "192.168.0.106:20001"

	//解析TCP地址及端口
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	//监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("start handler client.")
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("read error:", n)
			return
		}
		fmt.Println("read number is:", n)

		//write data，收到了
		n, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
		fmt.Println("write number is:", n)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
