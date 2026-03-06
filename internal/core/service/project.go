package service

import (
	"context"
	"fmt"

	"dev.harrysoler/todoweb/internal/core/entity"
	"dev.harrysoler/todoweb/internal/core/repository"
	"github.com/segmentio/ksuid"
)

type ProjectService struct {
	repository repository.ProjectRepository
}

func NewProjectService(repository repository.ProjectRepository) ProjectService {
	return ProjectService{
		repository: repository,
	}
}

func (service *ProjectService) Projects(ctx context.Context) ([]entity.Project, error) {
	return service.repository.Projects(ctx)
}

func (service *ProjectService) CreateProject(ctx context.Context, request entity.CreateProjectRequest) error {
	isProjectDuplicated, err := service.repository.IsProjectDuplicated(ctx, request.Name)

	if err != nil {
		return fmt.Errorf("create project: %w", err)
	}

	if isProjectDuplicated {
		return fmt.Errorf("create project (name %s): %w", request.Name, entity.ErrProjectAlreadyExist)
	}

	project, err := request.ToEntity(generateProjectId(), generateTodoId)

	if err != nil {
		return fmt.Errorf("create project: %w", err)
	}

	err = service.repository.SaveProject(ctx, project)

	if err != nil {
		return fmt.Errorf("create project: %w", err)
	}

	return nil
}

func generateTodoId() entity.TodoId {
	return entity.TodoId(ksuid.New().String())
}

func generateProjectId() entity.ProjectId {
	return entity.ProjectId(ksuid.New().String())
}
