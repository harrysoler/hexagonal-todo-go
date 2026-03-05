package entity

import "errors"

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

type CreateProjectRequest struct {
	Name  ProjectName
	Todos []CreateTodoRequest
}
