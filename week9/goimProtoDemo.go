package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9123")
	if err != nil {
		panic(err.Error())
	}

	go Serve(listener)

	// ugly, demo only
	fmt.Println("Wait for server to start")
	time.Sleep(time.Duration(2) * time.Second)

	// test for multiple clients request at the same time
	for i := 1; i < 3; i++ {
		msg := newProtoMsg(0, 4, int32(i), []byte("Hello from the outside~"))
		go Browse(msg)
	}

	input := bufio.NewReader(os.Stdin)
	// press enter to exit
	input.ReadString('\n')
}
