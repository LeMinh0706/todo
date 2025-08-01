package todo

import (
	"context"

	"github.com/LeMinh0706/todo/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddService struct {
	list *List
	proto.UnimplementedAddTodoServiceServer
}

// AddTodo implements proto.AddTodoServiceServer.
func (s *AddService) AddTodo(ctx context.Context, req *proto.AddTodoRequest) (*proto.AddTodoResponse, error) {
	if req.Todo.GetDescription() == "" {
		return nil, status.Error(codes.InvalidArgument, "description cannot be empty")
	}

	id, _ := uuid.NewRandom()

	todo := &Todo{
		ID:          id.String(),
		Description: req.Todo.GetDescription(),
	}

	s.list.Add(*todo)

	return &proto.AddTodoResponse{
		Id: todo.ID,
	}, nil
}

func NewAddService(list *List) proto.AddTodoServiceServer {
	return &AddService{
		list: list,
	}
}
