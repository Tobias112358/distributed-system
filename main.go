package main

import (
	"os"

	"github.com/Tobias112358/distributed-system/core"
)

func main() {
	nodeType := os.Args[1]
	println("*****")
	println(nodeType)
	println(os.Args[2])
	println("*****")
	switch nodeType {
	case "master":
		core.GetMasterNode().Start()
	case "worker":
		core.GetWorkerNode(os.Args[2]).Start()
	default:
		panic("invalid node type")
	}
}
