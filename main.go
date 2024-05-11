package main

import (
	"os"

	"github.com/Tobias112358/distributed-system/core"
)

func main() {
	nodeType := os.Args[0]
	println("*****")
	println(nodeType)
	println("*****")
	switch nodeType {
	case "master":
		core.GetMasterNode().Start()
	case "worker":
		core.GetWorkerNode().Start()
	default:
		panic("invalid node type")
	}
}
