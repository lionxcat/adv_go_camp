package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// http handle
func demoHandle(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "[%s] Hello %s!\r\n", time.Now(), req.RemoteAddr)
}

// handleSignal handle linux signals
// shutdown http server gracefully when SIGINT/SIGHUP/SIGTERM recieved
func handleSignal(servers ...*http.Server) {
	sigChan := make(chan os.Signal, 1)
	// register signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	// wait for signal
	c := <-sigChan
	fmt.Printf("received signal: %v, shuting down server...\r\n", c)
	// shuting down all servers with each a 3 seconds time bound
	for _, svr := range servers {
		fmt.Printf("server %s will shutdown in less than 3 seconds\r\n", svr.Addr)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		svr.Shutdown(ctx)
	}
}

func main() {
	signal.Ignore()

	// register http servers
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", demoHandle)
	server1 := &http.Server{
		Addr:    ":8080",
		Handler: mux1,
	}

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", demoHandle)
	server2 := &http.Server{
		Addr:    ":8081",
		Handler: mux2,
	}

	servers := [](*http.Server){server1, server2}

	// lanch http server with errgroup
	group, _ := errgroup.WithContext(context.Background())

	for _, svr := range servers {
		svr := svr
		group.Go(func() error {
			err := svr.ListenAndServe()
			// if server brokes, exit program
			if err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
			return err
		})
	}

	// lanch signal handler on the background
	go handleSignal(server1, server2)

	// wait for http servers to exit
	// if either server fails with error or a shutdown signal recieved
	if err := group.Wait(); err != nil {
		fmt.Printf("server exit with error: %s\r\n", err.Error())
	}
}
