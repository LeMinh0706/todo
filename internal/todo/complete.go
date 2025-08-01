package todo

import (
	"context"

	"github.com/LeMinh0706/todo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompleteService struct {
	list *List
	proto.UnimplementedCompleteTodoServiceServer
}

func (c *CompleteService) CompleteTodo(ctx context.Context, req *proto.CompleteTodoRequest) (*proto.CompleteTodoResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id cannot be empty")

	}
	todo := c.list.Get(req.Id)
	if todo == nil {
		return nil, status.Error(codes.NotFound, "todo not found")
	}

	c.list.Delete(req.Id)

	return &proto.CompleteTodoResponse{}, nil
}

func NewCompleteService(list *List) proto.CompleteTodoServiceServer {
	return &CompleteService{
		list: list,
	}
}
