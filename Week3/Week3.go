package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func demoHandle(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "[%s] Hello %s from %s!\r\n", time.Now(), req.RemoteAddr, req.RequestURI)
}

func startHttpServer(ctx context.Context, route, addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc(route, demoHandle)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("server %s exit with error %+v", server.Addr, err)
		}
	}()
	<-ctx.Done()
	fmt.Printf("server %s will shutdown in less than 10 seconds", addr)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	fmt.Printf("server %s shutdown gracefully", addr)
}

func main() {
	signal.Ignore()
	sigChan := make(chan os.Signal, 1)
	// register signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go startHttpServer(ctx, "/", ":8080")

	// wait for signals
	fmt.Printf("received signal: %v", <-sigChan)
	// gracefully shutdown http server
	fmt.Printf("shuting down server")
	cancel()

}
