package core

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerNode struct {
	conn *grpc.ClientConn  // grpc client connection
	c    NodeServiceClient // grpc client
}

func (n *WorkerNode) Init(ip string) (err error) {
	// connect to master node
	n.conn, err = grpc.Dial(ip+":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	// grpc client
	n.c = NewNodeServiceClient(n.conn)

	return nil
}

func (n *WorkerNode) Start() {
	// log
	fmt.Println("worker node started")

	// report status
	_, _ = n.c.ReportStatus(context.Background(), &Request{})

	// assign task
	stream, _ := n.c.AssignTask(context.Background(), &Request{})
	for {
		// receive command from master node
		res, err := stream.Recv()
		if err != nil {
			return
		}

		// log command
		fmt.Println("received command: ", res.Data)

		// execute command
		parts := strings.Split(res.Data, " ")
		if err := exec.Command(parts[0], parts[1:]...).Run(); err != nil {
			fmt.Println("doesnt work")
			fmt.Println(err)
		}
	}
}

var workerNode *WorkerNode

func GetWorkerNode(ip string) *WorkerNode {
	if workerNode == nil {
		// node
		workerNode = &WorkerNode{}

		// initialize node
		if err := workerNode.Init(ip); err != nil {
			panic(err)
		}
	}

	return workerNode
}
