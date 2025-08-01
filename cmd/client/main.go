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

	add := proto.NewAddTodoServiceClient(conn)
	get := proto.NewListTaskServiceClient(conn)

	_, err = add.AddTodo(ctx, &proto.AddTodoRequest{
		Todo: &proto.Todo{
			Description: "Wake up",
		},
	})

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Fatal("failed to call AddTodo:", err)
	}

	_, err = add.AddTodo(ctx, &proto.AddTodoRequest{
		Todo: &proto.Todo{
			Description: "Breakfast",
		},
	})

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Fatal("failed to call AddTodo:", err)
	}

	complete := proto.NewCompleteTodoServiceClient(conn)
	_, err = complete.CompleteTodo(ctx, &proto.CompleteTodoRequest{
		Id: "96222013-9283-4946-bdd5-683490182c81",
	})

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Fatal("failed to call CompleteTodo:", err)
	}

	list, err := get.ListTasks(ctx, &proto.ListTasksRequest{})

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("RPC failed with status code %s: %v", status.Code(), status.Message())
		}
		log.Fatal("failed to call list:", err)
	}

	log.Println("Response from server:", list)
}
