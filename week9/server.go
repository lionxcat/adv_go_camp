package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// tcp server
func Serve( /* ctx context.Context */ l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("[server] accept failed, err:%v\n", err)
			return
		}
		defer conn.Close()
		go Process(conn)
	}
}

// server side message process routine
func Process(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := readSock(reader)
		if err != nil {
			if err == io.EOF {
				fmt.Println("[server] read closed")
			} else {
				fmt.Printf("[server] read failed, err:%v\n", err)
			}
			break
		}
		fmt.Printf("\033[1;31;40m[client says] %s\033[0m\n", msg.String())

		// print client side message
		err = newProtoMsg(msg.ver, msg.op, msg.seq, []byte("You must've called a thousand times!")).writeSock(conn)
		if err != nil {
			fmt.Printf("[server] write failed, err:%v\n", err)
			break
		}
		//fmt.Println("server sent")
	}
}
