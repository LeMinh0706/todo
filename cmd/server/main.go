package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/LeMinh0706/todo/internal/streaming"
	"github.com/LeMinh0706/todo/internal/todo"
	"github.com/LeMinh0706/todo/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil && errors.Is(err, context.Canceled) {
		log.Fatalf("failed to run server: %v", err)
	}

}

func run(ctx context.Context) error {
	grpcServer := grpc.NewServer()

	list := &todo.List{}
	service := todo.NewService(list)
	stream := streaming.NewService()

	proto.RegisterTodoServiceServer(grpcServer, service)
	proto.RegisterStreamingServiceServer(grpcServer, stream)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}
		log.Println("gRPC server is running on port 50051")

		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve gRPC server: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		log.Println("shutting down gRPC server")
		grpcServer.GracefulStop()
		return nil
	})

	return g.Wait()
}
