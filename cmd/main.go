package main

import (
	"os"

	"github.com/GoPlayAndFun/Distributed-File-System/internal/lockservice"
	"github.com/GoPlayAndFun/Distributed-File-System/internal/node"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.New(os.Stdout).With()

	log := zerolog.New(os.Stdout).With().Logger().Level(zerolog.GlobalLevel())
	ls := lockservice.NewSimpleLockService(log)
	node.Start(ls)
}
