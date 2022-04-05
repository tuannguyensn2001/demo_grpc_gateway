package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	taskpb "grpc_test/proto/task"
	"grpc_test/src/services/task/transport"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("listen error", err)
	}

	s := grpc.NewServer()

	taskServer := tasktransport.NewTaskServer()

	taskpb.RegisterTaskServiceServer(s, taskServer)

	go func() {
		log.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8000", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to dial server", err)
	}

	gwmux := runtime.NewServeMux()

	err = taskpb.RegisterTaskServiceHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatal("Failed to register gateway", err)
	}

	gwServer := &http.Server{Addr: ":8081", Handler: gwmux}

	log.Println("listen on 8081")
	err = gwServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
