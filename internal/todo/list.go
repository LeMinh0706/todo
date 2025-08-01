package todo

import (
	"context"

	"github.com/LeMinh0706/todo/proto"
)

type GetService struct {
	proto.UnimplementedListTaskServiceServer
	list *List
}

// ListTasks implements proto.ListTaskServiceServer.
func (g *GetService) ListTasks(ctx context.Context, req *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	todos := g.list.GetAll()
	var protoTodos []*proto.Todo

	for _, todo := range todos {
		protoTodos = append(protoTodos, &proto.Todo{
			Id:          todo.ID,
			Description: todo.Description,
		})
	}

	return &proto.ListTasksResponse{Todos: protoTodos}, nil
}

func NewGetService(list *List) proto.ListTaskServiceServer {
	return &GetService{list: list}
}
