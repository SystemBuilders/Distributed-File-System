package main

import (
	"os"

	"github.com/GoPlayAndFun/Distributed-File-System/internal/lockserver"
)

func main() {
	shutdownSignal := make(chan os.Signal, 1)
	lockserver.StartServer(shutdownSignal)
}
