package main

import (
	"log"
	"net"

	"github.com/LeMinh0706/todo/internal/todo"
	"github.com/LeMinh0706/todo/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	list := &todo.List{}

	add := todo.NewAddService(list)
	complete := todo.NewCompleteService(list)
	get := todo.NewGetService(list)

	proto.RegisterAddTodoServiceServer(grpcServer, add)
	proto.RegisterCompleteTodoServiceServer(grpcServer, complete)
	proto.RegisterListTaskServiceServer(grpcServer, get)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	log.Println("gRPC server is running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
