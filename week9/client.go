package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// client side routine
func Browse(msg *protoMsg) {
	conn, err := net.Dial("tcp", "127.0.0.1:9123")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// test for one client requests multiple times
	for i := 0; i < 10; i++ {
		// add a version
		msg.ver++
		err = msg.writeSock(conn)
		if err != nil {
			fmt.Printf("[server] send failed, err:%v\n", err)
			return
		}
		//fmt.Println("client sent")
		reader := bufio.NewReader(conn)
		server_msg, err := readSock(reader)
		if err != nil {
			if err == io.EOF {
				fmt.Println("[server] read closed")
			} else {
				fmt.Printf("[server] read failed, err:%v\n", err)
			}
			return
		}

		// print server side response
		fmt.Printf("\033[1;32;40m[server says] %s\033[0m\n", server_msg.String())
	}
}
