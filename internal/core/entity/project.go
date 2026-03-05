package entity

import (
	"errors"
	"fmt"
)

var (
	ErrProjectAlreadyExist = errors.New("project already exist")
)

type ProjectId string
type ProjectName string

type Project struct {
	id    ProjectId
	name  ProjectName
	todos []Todo
}

// TODO: validation
func NewProject(id ProjectId, name ProjectName, todos []Todo) (Project, error) {
	return Project{
		id:    id,
		name:  name,
		todos: todos,
	}, nil
}

func (project Project) Name() ProjectName {
	return project.name
}

type CreateProjectRequest struct {
	Name  ProjectName
	Todos []CreateTodoRequest
}

func (request CreateProjectRequest) ToEntity(id ProjectId, generateTodoId func() TodoId) (Project, error) {
	todos := make([]Todo, len(request.Todos))

	for index, todoRequest := range request.Todos {
		todo, err := todoRequest.ToEntity(generateTodoId())

		if err != nil {
			return Project{}, fmt.Errorf("create todo: %w", err)
		}

		todos[index] = todo
	}

	return NewProject(id, request.Name, todos)
}
