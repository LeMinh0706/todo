package main

import (
	"context"
	"log"

	"github.com/LeMinh0706/todo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("failed to connect to server:", err)
	}
	defer conn.Close()
	client := proto.NewTodoServiceClient(conn)

	_, err = client.AddTodo(ctx, &proto.AddTodoRequest{Todo: &proto.Todo{Description: "Wake up"}})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Printf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Println("failed to call AddTodo:", err)
	}

	_, err = client.AddTodo(ctx, &proto.AddTodoRequest{Todo: &proto.Todo{Description: "Breakfast"}})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Printf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Println("failed to call AddTodo:", err)
	}

	_, err = client.CompleteTodo(ctx, &proto.CompleteTodoRequest{Id: "fcddf572-2468-46f7-95f4-11ae9ba918f4"})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Printf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Println("failed to call CompleteTodo:", err)
	}

	list, err := client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Printf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Println("failed to call ListTasks:", err)
	}

	log.Println("Todo List:", list.Todos)

	log.Println("Connected to gRPC server")
}
