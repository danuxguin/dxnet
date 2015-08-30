package main

import (
	"flag"
	"lanstonetech.com/server"
	"math/rand"
	"runtime"
	"time"
)

type RunProc func(cmd []string)
type ServerRunCmd struct {
	Run RunProc
}

var RunServerCommand map[string]ServerRunCmd

func init() {

	RunServerCommand = make(map[string]ServerRunCmd)

	RunServerCommand["server"] = ServerRunCmd{server.OnRun}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		return
	}

	server, find := RunServerCommand[args[0]]
	if !find {
		return
	}

	server.Run(args[1:])
}
