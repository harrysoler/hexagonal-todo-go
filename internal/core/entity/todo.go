package entity

type TodoId string
type TodoName string

type Todo struct {
	id   TodoId
	name TodoName
}

func NewTodo(id TodoId, name TodoName) (Todo, error) {
	return Todo{
		id:   id,
		name: name,
	}, nil
}

func (todo Todo) Name() TodoName {
	return todo.name
}

type CreateTodoRequest struct {
	Name TodoName
}

func (request CreateTodoRequest) ToEntity(id TodoId) (Todo, error) {
	return NewTodo(id, request.Name)
}
