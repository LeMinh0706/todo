package todo

import (
	"context"

	"github.com/LeMinh0706/todo/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedTodoServiceServer
	list *List
}

// AddTodo implements proto.TodoServiceServer.
func (s *Service) AddTodo(ctx context.Context, req *proto.AddTodoRequest) (*proto.AddTodoResponse, error) {
	if req.Todo.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "Description is required")
	}

	id, _ := uuid.NewRandom()
	req.Todo.Id = id.String()
	s.list.Add(Todo{
		ID:          id.String(),
		Description: req.Todo.Description,
	})

	return &proto.AddTodoResponse{Id: req.Todo.Id}, nil
}

// CompleteTodo implements proto.TodoServiceServer.
func (s *Service) CompleteTodo(ctx context.Context, req *proto.CompleteTodoRequest) (*proto.CompleteTodoResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is required")
	}
	todo := s.list.Get(req.Id)
	if todo == nil {
		return nil, status.Error(codes.NotFound, "Todo not found")
	}

	s.list.Delete(req.Id)
	return &proto.CompleteTodoResponse{}, nil
}

// ListTasks implements proto.TodoServiceServer.
func (s *Service) ListTasks(ctx context.Context, req *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	todos := s.list.GetAll()
	var protoTodos []*proto.Todo
	for _, todo := range todos {
		protoTodos = append(protoTodos, &proto.Todo{
			Id:          todo.ID,
			Description: todo.Description,
		})
	}

	return &proto.ListTasksResponse{Todos: protoTodos}, nil
}

func NewAddService(list *List) proto.TodoServiceServer {
	return &Service{list: list}
}
