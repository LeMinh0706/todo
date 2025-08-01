package todo

import "log"

type Todo struct {
	ID          string
	Description string
}
type List struct {
	Todos []Todo
}

func (l *List) Add(todo Todo) {
	l.Todos = append(l.Todos, todo)
}

func (l *List) Get(id string) *Todo {
	for _, todo := range l.Todos {
		if todo.ID == id {
			log.Println("Found todo:", todo)
			return &todo
		}
	}
	return nil
}

func (l *List) GetAll() []Todo {
	return l.Todos
}

func (l *List) Delete(id string) {
	for i, todo := range l.Todos {
		if todo.ID == id {
			l.Todos = append(l.Todos[:i], l.Todos[i+1:]...)
			break
		}
	}
}
