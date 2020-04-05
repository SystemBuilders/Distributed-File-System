package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/suraj44/Distributed-File-System/internal/lockserver"
)

func main() {
	service := new(lockserver.Service)
	port := flag.Int("port", 55550, "service port")
	ip := flag.String("ip", "0.0.0.0", "service ip")

	rpcServer := rpc.NewServer()
	err := rpcServer.Register(service)
	if err != nil {
		log.Fatalf("Error in registering the RPC server: %v\n", err)
	}

	server := &http.Server{
		Addr:    *ip + ":" + strconv.Itoa(*port),
		Handler: rpcServer,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("listen error:", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
	cancel()
}
